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

// Package client provides a client for the FISCO BCOS RPC API.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"errors"

	"github.com/KasperLiu/gobcos/common"
	"github.com/KasperLiu/gobcos/common/hexutil"
	"github.com/KasperLiu/gobcos/rpc"
	"github.com/KasperLiu/gobcos/core/types"
	"github.com/KasperLiu/gobcos/rlp"
)

// Client defines typed wrappers for the Ethereum RPC API. 
type Client struct {
	c       *rpc.Client
	groupID uint
}

// Dial connects a client to the given URL and groupID.
func Dial(rawurl string, groupID uint) (*Client, error) {
	return DialContext(context.Background(), rawurl, groupID)
}

// DialContext pass the context to the rpc client
func DialContext(ctx context.Context, rawurl string, groupID uint) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	client := NewClient(c, groupID)
	_, err = client.GetClientVersion(context.Background())
	if err != nil {
		return nil, fmt.Errorf("the RPC server does not support the FISCO BCOS RPC API: GroupID or URL is wrong")
	}

	return client, nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, groupID uint) *Client {
	return &Client{c: c, groupID: groupID}
}

// Close disconnects the rpc
func (gc *Client) Close() {
	gc.c.Close()
}

// ============================================== FISCO BCOS Blockchain Access ================================================

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (gc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	// ======================================== KasperLiu =========================================
	err := gc.c.CallContext(ctx, &result, "getCode", gc.groupID, account.String())
	return result, err
}


// Filters

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

// FilterLogs executes a filter query.
func (gc *Client) FilterLogs(ctx context.Context, q common.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = gc.c.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (gc *Client) SubscribeFilterLogs(ctx context.Context, q common.FilterQuery, ch chan<- types.Log) (common.Subscription, error) {
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	return gc.c.EthSubscribe(ctx, ch, "logs", arg)
}

func toFilterArg(q common.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{
		"address": q.Addresses,
		"topics":  q.Topics,
	}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

// Pending State

// PendingCodeAt returns the contract code of the given account in the pending state.
func (gc *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var result hexutil.Bytes
	err := gc.c.CallContext(ctx, &result, "getCode", gc.groupID, account.String())
	return result, err
}

// Contract Calling

type callResult struct {
	CurrentBlockNumber string `json:"currentBlockNumber"`
	Output             string `json:"output"`
	Status             string `json:"status"`
}

// CallContract invoke the call method of rpc api
func (gc *Client) CallContract(ctx context.Context, msg common.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	var cr *callResult
	err := gc.c.CallContext(ctx, &cr, "call", gc.groupID, toCallArg(msg))
	if err != nil {
		return nil, err
	}
	hex = common.FromHex(cr.Output)
	return hex, nil
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (gc *Client) PendingCallContract(ctx context.Context, msg common.CallMsg) ([]byte, error) {
	var hex hexutil.Bytes
	err := gc.c.CallContext(ctx, &hex, "call", gc.groupID, toCallArg(msg))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (gc *Client) SendTransaction(ctx context.Context, tx *types.RawTransaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		fmt.Printf("rlp encode tx error!")
		return err
	}
	return gc.c.CallContext(ctx, nil, "sendRawTransaction", gc.groupID, common.ToHex(data))
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (gc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := gc.c.CallContext(ctx, &r, "getTransactionReceipt", gc.groupID, txHash.Hex())
	if err == nil {
		if r == nil {
			return nil, errors.New("Transaction not found")
		}
	}
	return r, err
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

// GetGroupID returns the groupID of the client
func (gc *Client) GetGroupID() *big.Int {
	return big.NewInt(int64(gc.groupID))
}

// SetGroupID sets the groupID of the client
func (gc *Client) SetGroupID(newID uint) {
	gc.groupID = newID
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (gc *Client) GetClientVersion(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getClientVersion")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetChainID returns the Chain ID of the FISCO BCOS running on the nodes.
func (gc *Client) GetChainID(ctx context.Context) (*big.Int, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getClientVersion")
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
func (gc *Client) GetBlockNumber(ctx context.Context) ([]byte, error) {
	var raw string
	err := gc.c.CallContext(ctx, &raw, "getBlockNumber", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockLimit returns the blocklimit for current blocknumber
func (gc *Client) GetBlockLimit(ctx context.Context) (*big.Int, error) {
	defaultNumber := big.NewInt(500)
	var raw hexutil.Big
	err := gc.c.CallContext(ctx, &raw, "getBlockNumber", gc.groupID)
	if err != nil {
		return nil, err
	}
	return defaultNumber.Add(defaultNumber, (*big.Int)(&raw)), nil
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (gc *Client) GetPBFTView(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPbftView", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err

	// TODO
	// Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (gc *Client) GetSealerList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getSealerList", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (gc *Client) GetObserverList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getObserverList", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (gc *Client) GetConsensusStatus(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getConsensusStatus", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetSyncStatus returns the synchronization status of the group
func (gc *Client) GetSyncStatus(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getSyncStatus", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPeers returns the information of the connected peers
func (gc *Client) GetPeers(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPeers", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (gc *Client) GetGroupPeers(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getGroupPeers", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (gc *Client) GetNodeIDList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getNodeIDList", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetGroupList returns the groupID list that the node belongs to
func (gc *Client) GetGroupList(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getGroupList")
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockByHash returns the block information according to the given block hash
func (gc *Client) GetBlockByHash(ctx context.Context, bhash string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockByHash", gc.groupID, bhash, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (gc *Client) GetBlockByNumber(ctx context.Context, bnum string, includetx bool) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockByNumber", gc.groupID, bnum, includetx)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (gc *Client) GetBlockHashByNumber(ctx context.Context, bnum string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getBlockHashByNumber", gc.groupID, bnum)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (gc *Client) GetTransactionByHash(ctx context.Context, txhash string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionByHash", gc.groupID, txhash)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByBlockHashAndIndex returns the transaction information according to
// the given block hash and transaction index
func (gc *Client) GetTransactionByBlockHashAndIndex(ctx context.Context, bhash string, txindex string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionByBlockHashAndIndex", gc.groupID, bhash, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to
// the given block number and transaction index
func (gc *Client) GetTransactionByBlockNumberAndIndex(ctx context.Context, bnum string, txindex string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTransactionByBlockNumberAndIndex", gc.groupID, bnum, txindex)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (gc *Client) GetTransactionReceipt(ctx context.Context, txhash string) (*types.Receipt, error) {
	var raw *types.Receipt
	err := gc.c.CallContext(ctx, &raw, "getTransactionReceipt", gc.groupID, txhash)
	if err != nil {
		return nil, err
	}
	// js, err := json.MarshalIndent(raw, "", "\t")
	return raw, err
}

// GetContractAddress returns a contract address according to the transaction hash 
func (gc *Client) GetContractAddress(ctx context.Context, txhash string) (common.Address, error) {
	var raw interface{}
	var contractAddress common.Address
	err := gc.c.CallContext(ctx, &raw, "getTransactionReceipt", gc.groupID, txhash)
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
func (gc *Client) GetPendingTransactions(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPendingTransactions", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetPendingTxSize returns amount of the pending transactions
func (gc *Client) GetPendingTxSize(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getPendingTxSize", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetCode returns the contract code according to the contract address
func (gc *Client) GetCode(ctx context.Context, addr string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getCode", gc.groupID, addr)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (gc *Client) GetTotalTransactionCount(ctx context.Context) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getTotalTransactionCount", gc.groupID)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (gc *Client) GetSystemConfigByKey(ctx context.Context, findkey string) ([]byte, error) {
	var raw interface{}
	err := gc.c.CallContext(ctx, &raw, "getSystemConfigByKey", gc.groupID, findkey)
	if err != nil {
		return nil, err
	}
	js, err := json.MarshalIndent(raw, "", "\t")
	return js, err
}
