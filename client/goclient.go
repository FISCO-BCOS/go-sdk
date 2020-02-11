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
	"log"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/rpc"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	apiHandler        *rpc.APIHandler
	groupID           int
	chainID           int64
	compatibleVersion string
	auth              *bind.TransactOpts
	callOpts          *bind.CallOpts
	smCrypto          bool
}

const (
	//V210 is node version v2.1.0
	V210 = "2.1.0"
	//V220 is node version v2.2.0
	V220 = "2.2.0"
)

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
	client := Client{apiHandler: apiHandler, groupID: config.GroupID, compatibleVersion: compatibleVersion, chainID: config.ChainID, smCrypto: config.IsSMCrypto}
	if config.IsSMCrypto {
		client.auth = bind.NewSMCryptoTransactor(config.PrivateKey)
	} else {
		privateKey, err := crypto.HexToECDSA(config.PrivateKey)
		if err != nil {
			log.Fatal(err)
		}
		client.auth = bind.NewKeyedTransactor(privateKey)
	}
	client.auth.GasLimit = big.NewInt(30000000)
	client.callOpts = &bind.CallOpts{From: client.auth.From}
	return &client, nil
}

// Close disconnects the rpc
func (c *Client) Close() {
	c.apiHandler.Close()
}

// ============================================== FISCO BCOS Blockchain Access ================================================

// GetTransactOpts return *bind.TransactOpts
func (c *Client) GetTransactOpts() *bind.TransactOpts {
	return c.auth
}

// GetCallOpts return *bind.CallOpts
func (c *Client) GetCallOpts() *bind.CallOpts {
	return c.callOpts
}

// WaitMined is wrapper of bind.WaitMined
func (c *Client) WaitMined(tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(context.Background(), c, tx)
}

// SMCrypto returns true if use sm crypto
func (c *Client) SMCrypto() bool {
	return c.smCrypto
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (c *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return c.apiHandler.GetCode(ctx, c.groupID, account.String())
}

// Filters
func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

// FilterLogs executes a filter query.
func (c *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = c.apiHandler.CallContext(ctx, &result, "eth_getLogs", arg)
	return result, err
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (c *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	return c.apiHandler.EthSubscribe(ctx, ch, "logs", arg)
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
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
func (c *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return c.apiHandler.GetCode(ctx, c.groupID, account.String())
}

// Contract Calling

// CallContract invoke the call method of rpc api
func (c *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.apiHandler.Call(ctx, c.groupID, msg)
}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (c *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	return c.apiHandler.Call(ctx, c.groupID, msg)
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.apiHandler.SendRawTransaction(ctx, c.groupID, tx)
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (c *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return c.apiHandler.GetTransactionReceipt(ctx, c.groupID, txHash.Hex())
}

// GetGroupID returns the groupID of the client
func (c *Client) GetGroupID() *big.Int {
	return big.NewInt(int64(c.groupID))
}

// SetGroupID sets the groupID of the client
func (c *Client) SetGroupID(newID int) {
	c.groupID = newID
}

// GetClientVersion returns the version of FISCO BCOS running on the nodes.
func (c *Client) GetClientVersion(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetClientVersion(ctx)
}

// GetChainID returns the Chain ID of the FISCO BCOS running on the nodes.
func (c *Client) GetChainID(ctx context.Context) (*big.Int, error) {
	convertor := new(big.Int)
	var chainid *big.Int
	chainid = convertor.SetInt64(c.chainID)
	return chainid, nil
}

// GetBlockNumber returns the latest block height(hex format) on a given groupID.
func (c *Client) GetBlockNumber(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetBlockNumber(ctx, c.groupID)
}

// GetBlockLimit returns the blocklimit for current blocknumber
func (c *Client) GetBlockLimit(ctx context.Context) (*big.Int, error) {
	return c.apiHandler.GetBlockLimit(ctx, c.groupID)
}

// GetPBFTView returns the latest PBFT view(hex format) of the specific group and it will returns a wrong sentence
// if the consensus algorithm is not the PBFT.
func (c *Client) GetPBFTView(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetPBFTView(ctx, c.groupID)
	// TODO
	// Raft consensus
}

// GetSealerList returns the list of consensus nodes' ID according to the groupID
func (c *Client) GetSealerList(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetSealerList(ctx, c.groupID)
}

// GetObserverList returns the list of observer nodes' ID according to the groupID
func (c *Client) GetObserverList(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetObserverList(ctx, c.groupID)
}

// GetConsensusStatus returns the status information about the consensus algorithm on a specific groupID
func (c *Client) GetConsensusStatus(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetConsensusStatus(ctx, c.groupID)
}

// GetSyncStatus returns the synchronization status of the group
func (c *Client) GetSyncStatus(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetSyncStatus(ctx, c.groupID)
}

// GetPeers returns the information of the connected peers
func (c *Client) GetPeers(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetPeers(ctx, c.groupID)
}

// GetGroupPeers returns the nodes and the overser nodes list on a specific group
func (c *Client) GetGroupPeers(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetGroupPeers(ctx, c.groupID)
}

// GetNodeIDList returns the ID information of the connected peers and itself
func (c *Client) GetNodeIDList(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetNodeIDList(ctx, c.groupID)
}

// GetGroupList returns the groupID list that the node belongs to
func (c *Client) GetGroupList(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetGroupList(ctx)
}

// GetBlockByHash returns the block information according to the given block hash
func (c *Client) GetBlockByHash(ctx context.Context, bhash string, includetx bool) ([]byte, error) {
	return c.apiHandler.GetBlockByHash(ctx, c.groupID, bhash, includetx)
}

// GetBlockByNumber returns the block information according to the given block number(hex format)
func (c *Client) GetBlockByNumber(ctx context.Context, bnum string, includetx bool) ([]byte, error) {
	return c.apiHandler.GetBlockByNumber(ctx, c.groupID, bnum, includetx)
}

// GetBlockHashByNumber returns the block hash according to the given block number
func (c *Client) GetBlockHashByNumber(ctx context.Context, bnum string) ([]byte, error) {
	return c.apiHandler.GetBlockHashByNumber(ctx, c.groupID, bnum)

}

// GetTransactionByHash returns the transaction information according to the given transaction hash
func (c *Client) GetTransactionByHash(ctx context.Context, txhash string) ([]byte, error) {
	return c.apiHandler.GetTransactionByHash(ctx, c.groupID, txhash)
}

// GetTransactionByBlockHashAndIndex returns the transaction information according to
// the given block hash and transaction index
func (c *Client) GetTransactionByBlockHashAndIndex(ctx context.Context, bhash string, txindex string) ([]byte, error) {
	return c.apiHandler.GetTransactionByBlockHashAndIndex(ctx, c.groupID, bhash, txindex)
}

// GetTransactionByBlockNumberAndIndex returns the transaction information according to
// the given block number and transaction index
func (c *Client) GetTransactionByBlockNumberAndIndex(ctx context.Context, bnum string, txindex string) ([]byte, error) {
	return c.apiHandler.GetTransactionByBlockNumberAndIndex(ctx, c.groupID, bnum, txindex)
}

// GetTransactionReceipt returns the transaction receipt according to the given transaction hash
func (c *Client) GetTransactionReceipt(ctx context.Context, txhash string) (*types.Receipt, error) {
	return c.apiHandler.GetTransactionReceipt(ctx, c.groupID, txhash)
}

// GetContractAddress returns a contract address according to the transaction hash
func (c *Client) GetContractAddress(ctx context.Context, txhash string) (common.Address, error) {
	return c.apiHandler.GetContractAddress(ctx, c.groupID, txhash)
}

// GetPendingTransactions returns information of the pending transactions
func (c *Client) GetPendingTransactions(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetPendingTransactions(ctx, c.groupID)
}

// GetPendingTxSize returns amount of the pending transactions
func (c *Client) GetPendingTxSize(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetPendingTxSize(ctx, c.groupID)
}

// GetCode returns the contract code according to the contract address
func (c *Client) GetCode(ctx context.Context, addr string) ([]byte, error) {
	return c.apiHandler.GetCode(ctx, c.groupID, addr)
}

// GetTotalTransactionCount returns the totoal amount of transactions and the block height at present
func (c *Client) GetTotalTransactionCount(ctx context.Context) ([]byte, error) {
	return c.apiHandler.GetTotalTransactionCount(ctx, c.groupID)
}

// GetSystemConfigByKey returns value according to the key(only tx_count_limit, tx_gas_limit could work)
func (c *Client) GetSystemConfigByKey(ctx context.Context, configKey string) ([]byte, error) {
	return c.apiHandler.GetSystemConfigByKey(ctx, c.groupID, configKey)
}
