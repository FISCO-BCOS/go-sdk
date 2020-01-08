// @CopyRight:
// FISCO-BCOS go-sdk is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// FISCO-BCOS go-sdk is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with FISCO-BCOS go-sdk.  If not, see <http://www.gnu.org/licenses/>
// (c) 2016-2018 fisco-dev contributors.

// Package client provides a client for the FISCO BCOS RPC API.
package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/common/hexutil"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/rpc"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	apiHandler        *rpc.APIHandler
	groupID           int
	chainID           int64
	compatibleVersion string
}

// Dial connects a client to the given URL and groupID.
func Dial(config *conf.Config) (*Client, error) {
	return DialContext(context.Background(), config)
}

// DialContext pass the context to the rpc client
func DialContext(ctx context.Context, config *conf.Config) (*Client, error) {
	var c *rpc.Connection
	var err error
	if config.IsHTTP {
		c, err = rpc.DialContextHTTP(config.NodeURL)
	} else {
		c, err = rpc.DialContextChannel(config.NodeURL, config.CAFile, config.Cert, config.Key, config.GroupID)
	}
	if err != nil {
		return nil, err
	}
	apiHandler := rpc.NewAPIHandler(c)
	var response []byte
	response, err = apiHandler.GetClientVersion(ctx)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	var raw interface{}
	json.Unmarshal(response, &raw)
	m, ok := raw.(map[string]interface{})
	if ok != true {
		return nil, errors.New("parse response json to map error")
	}
	var compatibleVersion string
	compatibleVersion, ok = m["Supported Version"].(string)
	if ok != true {
		return nil, errors.New("Json respond does not contains the key : Supported Version")
	}
	var nodeChainID int64
	nodeChainID, err = strconv.ParseInt(m["Chain Id"].(string), 10, 64)
	if ok != true {
		return nil, errors.New("Json respond does not contains the key : Chain Id")
	}
	if config.ChainID != nodeChainID {
		return nil, errors.New("The chain ID of node is " + fmt.Sprint(nodeChainID) + ", but configuration is " + fmt.Sprint(config.ChainID))
	}
	client := Client{apiHandler: apiHandler, groupID: config.GroupID, compatibleVersion: compatibleVersion, chainID: config.ChainID}
	return &client, nil
}

// Close disconnects the rpc
func (gc *Client) Close() {
	gc.apiHandler.Close()
}

// ============================================== FISCO BCOS Blockchain Access ================================================

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (gc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return gc.apiHandler.GetCode(ctx, gc.groupID, account.String())
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
	err = gc.apiHandler.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (gc *Client) SubscribeFilterLogs(ctx context.Context, q common.FilterQuery, ch chan<- types.Log) (common.Subscription, error) {
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	return gc.apiHandler.EthSubscribe(ctx, ch, "logs", arg)
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
	return gc.apiHandler.GetCode(ctx, gc.groupID, account.String())
}

// Contract Calling

// CallContract invoke the call method of rpc api
func (gc *Client) CallContract(ctx context.Context, msg common.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return gc.apiHandler.Call(ctx, gc.groupID, msg)
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (gc *Client) PendingCallContract(ctx context.Context, msg common.CallMsg) ([]byte, error) {
	return gc.apiHandler.Call(ctx, gc.groupID, msg)
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (gc *Client) SendTransaction(ctx context.Context, tx *types.RawTransaction) error {
	return gc.apiHandler.SendRawTransaction(ctx, gc.groupID, tx)
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (gc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return gc.apiHandler.GetTransactionReceipt(ctx, gc.groupID, txHash.Hex())
}

// GetGroupID returns the groupID of the client
func (gc *Client) GetGroupID() *big.Int {
	return big.NewInt(int64(gc.groupID))
}

// SetGroupID sets the groupID of the client
func (gc *Client) SetGroupID(newID int) {
	gc.groupID = newID
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (gc *Client) GetClientVersion(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetClientVersion(ctx)
}

// GetChainID returns the Chain ID of the FISCO BCOS running on the nodes.
func (gc *Client) GetChainID(ctx context.Context) (*big.Int, error) {
	convertor := new(big.Int)
	var chainid *big.Int
	chainid = convertor.SetInt64(gc.chainID)
	return chainid, nil
}

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (gc *Client) GetBlockNumber(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetBlockNumber(ctx, gc.groupID)
}

// GetBlockLimit returns the blocklimit for current blocknumber
func (gc *Client) GetBlockLimit(ctx context.Context) (*big.Int, error) {
	return gc.apiHandler.GetBlockLimit(ctx, gc.groupID)
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (gc *Client) GetPBFTView(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetPBFTView(ctx, gc.groupID)
	// TODO
	// Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (gc *Client) GetSealerList(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetSealerList(ctx, gc.groupID)
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (gc *Client) GetObserverList(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetObserverList(ctx, gc.groupID)
}

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (gc *Client) GetConsensusStatus(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetConsensusStatus(ctx, gc.groupID)
}

// GetSyncStatus returns the synchronization status of the group
func (gc *Client) GetSyncStatus(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetSyncStatus(ctx, gc.groupID)
}

// GetPeers returns the information of the connected peers
func (gc *Client) GetPeers(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetPeers(ctx, gc.groupID)
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (gc *Client) GetGroupPeers(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetGroupPeers(ctx, gc.groupID)
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (gc *Client) GetNodeIDList(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetNodeIDList(ctx, gc.groupID)
}

// GetGroupList returns the groupID list that the node belongs to
func (gc *Client) GetGroupList(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetGroupList(ctx)
}

// GetBlockByHash returns the block information according to the given block hash
func (gc *Client) GetBlockByHash(ctx context.Context, bhash string, includetx bool) ([]byte, error) {
	return gc.apiHandler.GetBlockByHash(ctx, gc.groupID, bhash, includetx)
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (gc *Client) GetBlockByNumber(ctx context.Context, bnum string, includetx bool) ([]byte, error) {
	return gc.apiHandler.GetBlockByNumber(ctx, gc.groupID, bnum, includetx)
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (gc *Client) GetBlockHashByNumber(ctx context.Context, bnum string) ([]byte, error) {
	return gc.apiHandler.GetBlockHashByNumber(ctx, gc.groupID, bnum)

}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (gc *Client) GetTransactionByHash(ctx context.Context, txhash string) ([]byte, error) {
	return gc.apiHandler.GetTransactionByHash(ctx, gc.groupID, txhash)
}

// GetTransactionByBlockHashAndIndex returns the transaction information according to
// the given block hash and transaction index
func (gc *Client) GetTransactionByBlockHashAndIndex(ctx context.Context, bhash string, txindex string) ([]byte, error) {
	return gc.apiHandler.GetTransactionByBlockHashAndIndex(ctx, gc.groupID, bhash, txindex)
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to
// the given block number and transaction index
func (gc *Client) GetTransactionByBlockNumberAndIndex(ctx context.Context, bnum string, txindex string) ([]byte, error) {
	return gc.apiHandler.GetTransactionByBlockNumberAndIndex(ctx, gc.groupID, bnum, txindex)
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (gc *Client) GetTransactionReceipt(ctx context.Context, txhash string) (*types.Receipt, error) {
	return gc.apiHandler.GetTransactionReceipt(ctx, gc.groupID, txhash)
}

// GetContractAddress returns a contract address according to the transaction hash
func (gc *Client) GetContractAddress(ctx context.Context, txhash string) (common.Address, error) {
	return gc.apiHandler.GetContractAddress(ctx, gc.groupID, txhash)
}

// GetPendingTransactions returns information of the pending transactions
func (gc *Client) GetPendingTransactions(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetPendingTransactions(ctx, gc.groupID)
}

// GetPendingTxSize returns amount of the pending transactions
func (gc *Client) GetPendingTxSize(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetPendingTxSize(ctx, gc.groupID)
}

// GetCode returns the contract code according to the contract address
func (gc *Client) GetCode(ctx context.Context, addr string) ([]byte, error) {
	return gc.apiHandler.GetCode(ctx, gc.groupID, addr)
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (gc *Client) GetTotalTransactionCount(ctx context.Context) ([]byte, error) {
	return gc.apiHandler.GetTotalTransactionCount(ctx, gc.groupID)
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (gc *Client) GetSystemConfigByKey(ctx context.Context, configKey string) ([]byte, error) {
	return gc.apiHandler.GetSystemConfigByKey(ctx, gc.groupID, configKey)
}
