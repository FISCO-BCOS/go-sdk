// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package rpc provides a client for the FISCO BCOS RPC API.
package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/common/hexutil"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/rlp"
)

// APIHandler defines typed wrappers for the Ethereum RPC API.
// TODO: make client interface
type APIHandler struct {
	*Client
}

// NewAPIHandler create a new API handler
func NewAPIHandler(c *Client) *APIHandler {
	apiHandler := APIHandler{c}
	return &apiHandler
}

func toCallArg(msg common.CallMsg) interface{} {
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
func (api *APIHandler) Call(ctx context.Context, groupID uint, msg common.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	var cr *callResult
	err := api.CallContext(ctx, &cr, "call", groupID, toCallArg(msg))
	if err != nil {
		return nil, err
	}
	hex = common.FromHex(cr.Output)
	return hex, nil
}

// SendRawTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (api *APIHandler) SendRawTransaction(ctx context.Context, groupID uint, tx *types.RawTransaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		fmt.Printf("rlp encode tx error!")
		return err
	}
	return api.CallContext(ctx, nil, "sendRawTransaction", groupID, common.ToHex(data))
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (api *APIHandler) TransactionReceipt(ctx context.Context, groupID uint, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := api.CallContext(ctx, &r, "getTransactionReceipt", groupID, txHash.Hex())
	if err == nil {
		if r == nil {
			return nil, errors.New("Transaction not found")
		}
	}
	return r, err
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (api *APIHandler) GetClientVersion(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getClientVersion")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
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
	if ok != true {
		return nil, errors.New("GetChainID Json respond does not satisfy the type assertion: map[string]interface{}")
	}
	var temp interface{}
	temp, ok = m["Chain Id"]
	if ok != true {
		return nil, errors.New("Json respond does not contains the key : Chain Id")
	}
	var strChainid string
	strChainid, ok = temp.(string)
	if ok != true {
		return nil, errors.New("type assertion for Chain Id is wrong: not a string")
	}
	convertor := new(big.Int)
	var chainid *big.Int
	chainid, ok = convertor.SetString(strChainid, 10)
	if ok != true {
		return nil, errors.New("big.Int.SetString(): error for Chain Id")
	}
	return chainid, nil
}

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (api *APIHandler) GetBlockNumber(ctx context.Context, groupID uint) ([]byte, error) {
	var raw string
	err := api.CallContext(ctx, &raw, "getBlockNumber", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockLimit returns the blocklimit for current blocknumber
func (api *APIHandler) GetBlockLimit(ctx context.Context, groupID uint) (*big.Int, error) {
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
func (api *APIHandler) GetPBFTView(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPbftView", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err

	// TODO
	// Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (api *APIHandler) GetSealerList(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getSealerList", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (api *APIHandler) GetObserverList(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getObserverList", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (api *APIHandler) GetConsensusStatus(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getConsensusStatus", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetSyncStatus returns the synchronization status of the group
func (api *APIHandler) GetSyncStatus(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getSyncStatus", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPeers returns the information of the connected peers
func (api *APIHandler) GetPeers(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPeers", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (api *APIHandler) GetGroupPeers(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getGroupPeers", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (api *APIHandler) GetNodeIDList(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getNodeIDList", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetGroupList returns the groupID list that the node belongs to
func (api *APIHandler) GetGroupList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getGroupList")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockByHash returns the block information according to the given block hash
func (api *APIHandler) GetBlockByHash(ctx context.Context, groupID uint, bhash string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getBlockByHash", groupID, bhash, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (api *APIHandler) GetBlockByNumber(ctx context.Context, groupID uint, bnum string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getBlockByNumber", groupID, bnum, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (api *APIHandler) GetBlockHashByNumber(ctx context.Context, groupID uint, bnum string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getBlockHashByNumber", groupID, bnum)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (api *APIHandler) GetTransactionByHash(ctx context.Context, groupID uint, txhash string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTransactionByHash", groupID, txhash)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByBlockHashAndIndex returns the transaction information according to
// the given block hash and transaction index
func (api *APIHandler) GetTransactionByBlockHashAndIndex(ctx context.Context, groupID uint, bhash string, txindex string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTransactionByBlockHashAndIndex", groupID, bhash, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to
// the given block number and transaction index
func (api *APIHandler) GetTransactionByBlockNumberAndIndex(ctx context.Context, groupID uint, bnum string, txindex string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTransactionByBlockNumberAndIndex", groupID, bnum, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (api *APIHandler) GetTransactionReceipt(ctx context.Context, groupID uint, txhash string) (*types.Receipt, error) {
	var raw *types.Receipt
	err := api.CallContext(ctx, &raw, "getTransactionReceipt", groupID, txhash)
	if err != nil {
		return nil, err
	}
	// js, err := json.MarshalIndent(raw, "", "\t")
	return raw, err
}

// GetContractAddress returns a contract address according to the transaction hash
func (api *APIHandler) GetContractAddress(ctx context.Context, groupID uint, txhash string) (common.Address, error) {
	var raw interface{}
	var contractAddress common.Address
	err := api.CallContext(ctx, &raw, "getTransactionReceipt", groupID, txhash)
	if err != nil {
		return contractAddress, err
	}
	m, ok := raw.(map[string]interface{})
	if ok != true {
		return contractAddress, errors.New("GetContractAddress Json respond does not satisfy the type assertion: map[string]interface{}")
	}
	var temp interface{}
	temp, ok = m["contractAddress"]
	if ok != true {
		return contractAddress, errors.New("Json respond does not contains the key : contractAddress")
	}
	var strContractAddress string
	strContractAddress, ok = temp.(string)
	if ok != true {
		return contractAddress, errors.New("type assertion for Chain Id is wrong: not a string")
	}
	return common.HexToAddress(strContractAddress), nil
}

// GetPendingTransactions returns information of the pending transactions
func (api *APIHandler) GetPendingTransactions(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPendingTransactions", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPendingTxSize returns amount of the pending transactions
func (api *APIHandler) GetPendingTxSize(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getPendingTxSize", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetCode returns the contract code according to the contract address
func (api *APIHandler) GetCode(ctx context.Context, groupID uint, addr string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getCode", groupID, addr)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (api *APIHandler) GetTotalTransactionCount(ctx context.Context, groupID uint) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getTotalTransactionCount", groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (api *APIHandler) GetSystemConfigByKey(ctx context.Context, groupID uint, configKey string) ([]byte, error) {
	var raw interface{}
	err := api.CallContext(ctx, &raw, "getSystemConfigByKey", groupID, configKey)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}
