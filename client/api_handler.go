// Copyright FISCO-BCOS go-sdk
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package client provides a client for the FISCO BCOS RPC API.
package client

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/FISCO-BCOS/go-sdk/conn"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

// APIHandler defines typed wrappers for the FISCO BCOS RPC API.
type APIHandler struct {
	*conn.Connection
}

const (
	indent = "  "
)

// NewAPIHandler create a new API handler
func NewAPIHandler(c *conn.Connection) *APIHandler {
	apiHandler := APIHandler{c}
	return &apiHandler
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From.String(),
		"to":   msg.To.String(),
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data).String()
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value).String()
	}

	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}

	return arg
}

type callResult struct {
	CurrentBlockNumber string `json:"currentBlockNumber"`
	Output             string `json:"output"`
	Status             string `json:"status"`
}

// Call invoke the call method of rpc api
func (api *APIHandler) Call(ctx context.Context, groupID int, msg ethereum.CallMsg) ([]byte, error) {
	var hexBytes hexutil.Bytes
	var cr *callResult
	err := api.CallContext(ctx, &cr, "call", groupID, toCallArg(msg))
	if err != nil {
		return nil, err
	}
	if cr.Status != "0x0" {
		var errorMessage string
		if len(cr.Output) >= 138 {
			outputBytes, err := hex.DecodeString(cr.Output[2:])
			if err != nil {
				return nil, fmt.Errorf("call error of status %s, hex.DecodeString failed", cr.Status)
			}
			errorMessage = string(outputBytes[68:])
		}
		return nil, fmt.Errorf("call error of status %s, %v", cr.Status, errorMessage)
	}
	hexBytes = common.FromHex(cr.Output)
	return hexBytes, nil
}

// SendRawTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (api *APIHandler) SendRawTransaction(ctx context.Context, groupID int, tx *types.Transaction) (*types.Receipt, error) {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		fmt.Printf("rlp encode tx error!")
		return nil, err
	}
	var receipt *types.Receipt
	if api.IsHTTP() {
		err = api.CallContext(ctx, nil, "sendRawTransaction", groupID, hexutil.Encode(data))
		if err != nil {
			return nil, err
		}

		// timer to wait transaction on-chain
		queryTicker := time.NewTicker(time.Second)
		defer queryTicker.Stop()
		logger := log.New("hash", tx.Hash())
		for {
			receipt, err := api.GetTransactionReceipt(ctx, groupID, tx.Hash().Hex())
			if receipt != nil {
				return receipt, nil
			}
			if err != nil {
				logger.Trace("Receipt retrieval failed", "err", err)
			} else {
				logger.Trace("Transaction not yet mined")
			}
			// Wait for the next round.
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-queryTicker.C:
			}
		}
	} else {
		var anonymityReceipt = &struct {
			types.Receipt
			Status string `json:"status"`
		}{}
		err = api.CallContext(ctx, anonymityReceipt, "sendRawTransaction", groupID, hexutil.Encode(data))
		if err != nil {
			return nil, err
		}
		status, err := strconv.ParseInt(anonymityReceipt.Status[2:], 16, 32)
		if err != nil {
			return nil, fmt.Errorf("SendRawTransaction failed, strconv.ParseInt err: " + fmt.Sprint(err))
		}
		receipt = &anonymityReceipt.Receipt
		receipt.Status = int(status)
		return receipt, nil
	}
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (api *APIHandler) TransactionReceipt(ctx context.Context, groupID int, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := api.CallContext(ctx, &r, "getTransactionReceipt", groupID, txHash.Hex())
	if err == nil && r == nil {
		return nil, fmt.Errorf("TransactionReceipt failed, transaction not found, txHash is: %v, receipt is: %+v", txHash.Hex(), r)
	}
	return r, err
}

func (api *APIHandler) SubscribeTopic(topic string, handler func([]byte)) error {
	return api.Connection.SubscribeTopic(topic, handler)
}

func (api *APIHandler) SubscribeAuthTopic(topic string, privateKey *ecdsa.PrivateKey, handler func([]byte)) error {
	return api.Connection.SubscribeAuthTopic(topic, privateKey, handler)
}

func (api *APIHandler) PublishAuthTopic(topic string, publicKey []*ecdsa.PublicKey, handler func([]byte)) error {
	return api.Connection.PublishAuthTopic(topic, publicKey, handler)
}

func (api *APIHandler) UnsubscribeTopic(topic string) error {
	return api.Connection.UnsubscribeTopic(topic)
}

func (api *APIHandler) UnsubscribeAuthTopic(topic string) error {
	return api.Connection.UnsubscribeAuthTopic(topic)
}

func (api *APIHandler) PushTopicDataRandom(topic string, data []byte) error {
	return api.Connection.PushTopicDataRandom(topic, data)
}

func (api *APIHandler) PushTopicDataToALL(topic string, data []byte) error {
	return api.Connection.PushTopicDataToALL(topic, data)
}

func (api *APIHandler) PushAuthTopicDataRandom(topic string, data []byte) error {
	return api.Connection.PushAuthTopicDataRandom(topic, data)
}

func (api *APIHandler) PushAuthTopicDataToALL(topic string, data []byte) error {
	return api.Connection.PushAuthTopicDataToALL(topic, data)
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (api *APIHandler) GetClientVersion(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getClientVersion")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetChainID returns the Chain ID of the FISCO BCOS running on the nodes.
func (api *APIHandler) GetChainID(ctx context.Context) (*big.Int, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getClientVersion")
	if err != nil {
		return nil, err
	}

	m, ok := raw.(map[string]interface{})
	if !ok {
		return nil, errors.New("GetChainID Json respond does not satisfy the type assertion: map[string]interface{}")
	}
	var temp interface{}
	temp, ok = m["Chain Id"]
	if !ok {
		return nil, errors.New("JSON respond does not contains the key : Chain Id")
	}
	var strChainid string
	strChainid, ok = temp.(string)
	if !ok {
		return nil, errors.New("type assertion for Chain Id is wrong: not a string")
	}
	convertor := new(big.Int)
	var chainid *big.Int
	chainid, ok = convertor.SetString(strChainid, 10)
	if !ok {
		return nil, errors.New("big.Int.SetString(): error for Chain Id")
	}
	return chainid, nil
}

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (api *APIHandler) GetBlockNumber(ctx context.Context, groupID int) ([]byte, error) {
	var raw string
	err := api.CallContext(ctx, &raw, "getBlockNumber", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetBlockLimit returns the blocklimit for current blocknumber
func (api *APIHandler) GetBlockLimit(ctx context.Context, groupID int) (*big.Int, error) {
	defaultNumber := big.NewInt(500)
	var raw hexutil.Big
	err := api.CallContext(ctx, &raw, "getBlockNumber", groupID)
	if err != nil {
		return nil, err
	}
	return defaultNumber.Add(defaultNumber, (*big.Int)(&raw)), nil
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (api *APIHandler) GetPBFTView(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPbftView", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err

	// TODO
	// Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (api *APIHandler) GetSealerList(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getSealerList", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (api *APIHandler) GetObserverList(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getObserverList", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (api *APIHandler) GetConsensusStatus(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getConsensusStatus", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetSyncStatus returns the synchronization status of the group
func (api *APIHandler) GetSyncStatus(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getSyncStatus", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetPeers returns the information of the connected peers
func (api *APIHandler) GetPeers(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPeers", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (api *APIHandler) GetGroupPeers(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getGroupPeers", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (api *APIHandler) GetNodeIDList(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getNodeIDList", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetGroupList returns the groupID list that the node belongs to
func (api *APIHandler) GetGroupList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getGroupList")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetBlockByHash returns the block information according to the given block hash
func (api *APIHandler) GetBlockByHash(ctx context.Context, groupID int, bhash string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getBlockByHash", groupID, bhash, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (api *APIHandler) GetBlockByNumber(ctx context.Context, groupID int, bnum string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getBlockByNumber", groupID, bnum, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (api *APIHandler) GetBlockHashByNumber(ctx context.Context, groupID int, bnum string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getBlockHashByNumber", groupID, bnum)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (api *APIHandler) GetTransactionByHash(ctx context.Context, groupID int, txhash string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTransactionByHash", groupID, txhash)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetTransactionByBlockHashAndIndex returns the transaction information according to
// the given block hash and transaction index
func (api *APIHandler) GetTransactionByBlockHashAndIndex(ctx context.Context, groupID int, bhash string, txindex string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTransactionByBlockHashAndIndex", groupID, bhash, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to
// the given block number and transaction index
func (api *APIHandler) GetTransactionByBlockNumberAndIndex(ctx context.Context, groupID int, bnum string, txindex string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTransactionByBlockNumberAndIndex", groupID, bnum, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (api *APIHandler) GetTransactionReceipt(ctx context.Context, groupID int, txhash string) (*types.Receipt, error) {
	var raw *types.Receipt
	var anonymityReceipt = &struct {
		types.Receipt
		Status string `json:"status"`
	}{}
	err := api.CallContext(ctx, anonymityReceipt, "getTransactionReceipt", groupID, txhash)
	if err != nil {
		return nil, err
	}
	if len(anonymityReceipt.Status) < 2 {
		return nil, fmt.Errorf("transaction %v is not on-chain", txhash)
	}
	status, err := strconv.ParseInt(anonymityReceipt.Status[2:], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionReceipt failed, strconv.ParseInt err: " + fmt.Sprint(err))
	}
	raw = &anonymityReceipt.Receipt
	raw.Status = int(status)
	return raw, err
}

// GetContractAddress returns a contract address according to the transaction hash
func (api *APIHandler) GetContractAddress(ctx context.Context, groupID int, txhash string) (common.Address, error) {
	var raw interface{}
	var contractAddress common.Address
	err := api.CallContext(ctx, &raw, "getTransactionReceipt", groupID, txhash)
	if err != nil {
		return contractAddress, err
	}
	m, ok := raw.(map[string]interface{})
	if !ok {
		return contractAddress, fmt.Errorf("GetContractAddress Json respond does not satisfy the type assertion: map[string]interface{}, %+v", raw)
	}
	var temp interface{}
	temp, ok = m["contractAddress"]
	if !ok {
		return contractAddress, errors.New("Json respond does not contains the key : contractAddress")
	}
	var strContractAddress string
	strContractAddress, ok = temp.(string)
	if !ok {
		return contractAddress, errors.New("type assertion for Chain Id is wrong: not a string")
	}
	return common.HexToAddress(strContractAddress), nil
}

// GetPendingTransactions returns information of the pending transactions
func (api *APIHandler) GetPendingTransactions(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPendingTransactions", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetPendingTxSize returns amount of the pending transactions
func (api *APIHandler) GetPendingTxSize(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPendingTxSize", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetCode returns the contract code according to the contract address
func (api *APIHandler) GetCode(ctx context.Context, groupID int, addr string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getCode", groupID, addr)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (api *APIHandler) GetTotalTransactionCount(ctx context.Context, groupID int) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTotalTransactionCount", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (api *APIHandler) GetSystemConfigByKey(ctx context.Context, groupID int, configKey string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getSystemConfigByKey", groupID, configKey)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", indent)
	return js, err
}
