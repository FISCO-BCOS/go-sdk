/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Package commandline is implement of console
package commandline

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/precompiled/config"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

var info = ", you can type console help for more information"

// commands
// var bashCompletionCmd = &cobra.Command{
// 	Use:   "bashCompletion",
// 	Short: "Generates bash completion scripts",
// 	Long: `A script "console.sh" will get you completions of the console commands.
// Copy it to

//     /etc/bash_completion.d/

// as described here:

//     https://debian-administration.org/article/316/An_introduction_to_bash_completion_part_1

// and reset your terminal to use autocompletion.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		rootCmd.GenBashCompletionFile("console.sh");
// 		fmt.Println("console.sh created on your current diretory successfully.")
// 	},
// }

// var zshCompletionCmd = &cobra.Command{
// 	Use:   "zshCompletion",
// 	Short: "Generates zsh completion scripts",
// 	Long: `A script "console.zsh" will get you completions of the console commands.
// The recommended way to install this script is to copy to '~/.zsh/_console', and
// then add the following to your ~/.zshrc file:

//     fpath=(~/.zsh $fpath)

// as described here:

//     https://debian-administration.org/article/316/An_introduction_to_bash_completion_part_1

// and reset your terminal to use autocompletion.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		rootCmd.GenZshCompletionFile("_console");
// 		fmt.Println("zsh file _console had created on your current diretory successfully.")
// 	},
// }

// =========== account ==========
var newAccountCmd = &cobra.Command{
	Use:   "newAccount",
	Short: "Create a new account",
	Long:  `Create a new account and save the results to ./bin/account/yourAccountName.keystore in encrypted form.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		clientVer, err := RPC.GetClientVersion(context.Background())
		if err != nil {
			fmt.Printf("client version not found: %v\n", err)
			return
		}
		fmt.Printf("Client Version: \n%s\n", clientVer)
	},
}

// ======= node =======

var getClientVersionCmd = &cobra.Command{
	Use:   "getClientVersion",
	Short: "                                 Get the blockchain version",
	Long:  `Returns the specific FISCO BCOS version that runs on the node you connected.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		clientVer, err := RPC.GetClientVersion(context.Background())
		if err != nil {
			fmt.Printf("client version not found: %v\n", err)
			return
		}
		fmt.Printf("Client Version: \n%s\n", clientVer)
	},
}

var getGroupIDCmd = &cobra.Command{
	Use:   "getGroupID",
	Short: "                                 Get the group ID of the client",
	Long:  `Returns the group ID that the console had connected to.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		groupID := RPC.GetGroupID()
		fmt.Printf("Group ID: \n%s\n", groupID)
	},
}

var getBlockNumberCmd = &cobra.Command{
	Use:   "getBlockNumber",
	Short: "                                 Get the latest block height of the blockchain",
	Long: `Returns the latest block height in the specified group.
The block height is encoded in hex`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		blockNumber, err := RPC.GetBlockNumber(context.Background())
		if err != nil {
			fmt.Printf("block number not found: %v\n", err)
			return
		}
		fmt.Printf("blocknumber: \n    hex: %s\n", blockNumber)
		strNum := string(blockNumber)
		bnum, err := toDecimal(strNum[3 : len(strNum)-1])
		if err != nil {
			fmt.Println("The block Number is tot a valid hex string")
			return
		}
		fmt.Println("decimal: ", bnum)
	},
}

var getPbftViewCmd = &cobra.Command{
	Use:   "getPbftView",
	Short: "                                 Get the latest PBFT view(PBFT consensus only)",
	Long: `Returns the latest PBFT view in the specified group where the node is located.
The PBFT view is encoded in hex`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		pbft, err := RPC.GetPBFTView(context.Background())
		if err != nil {
			fmt.Printf("PBFT view not found: %v\n", err)
			return
		}
		fmt.Printf("PBFT view: \n%s\n", pbft)
	},
}

var getSealerListCmd = &cobra.Command{
	Use:   "getSealerList",
	Short: "                                 Get the sealers' ID list",
	Long:  `Returns an ID list of the sealer nodes within the specified group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		sealerList, err := RPC.GetSealerList(context.Background())
		if err != nil {
			fmt.Printf("sealer list not found: %v\n", err)
			return
		}
		fmt.Printf("Sealer List: \n%s\n", sealerList)
	},
}

var getObserverListCmd = &cobra.Command{
	Use:   "getObserverList",
	Short: "                                 Get the observers' ID list",
	Long:  `Returns an ID list of observer nodes within the specified group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		observerList, err := RPC.GetObserverList(context.Background())
		if err != nil {
			fmt.Printf("observer list not found: %v\n", err)
			return
		}
		fmt.Printf("Observer List: \n%s\n", observerList)
	},
}

var getConsensusStatusCmd = &cobra.Command{
	Use:   "getConsensusStatus",
	Short: "                                 Get consensus status of nodes",
	Long:  `Returns consensus status information within the specified group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		consensusStatus, err := RPC.GetConsensusStatus(context.Background())
		if err != nil {
			fmt.Printf("consensus status not found: %v\n", err)
			return
		}
		fmt.Printf("Consensus Status: \n%s\n", consensusStatus)
	},
}

var getSyncStatusCmd = &cobra.Command{
	Use:   "getSyncStatus",
	Short: "                                 Get synchronization status of nodes",
	Long:  `Returns synchronization status information within the specified group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		syncStatus, err := RPC.GetSyncStatus(context.Background())
		if err != nil {
			fmt.Printf("synchronization status not found: %v\n", err)
			return
		}
		fmt.Printf("Synchronization Status: \n%s\n", syncStatus)
	},
}

var getPeersCmd = &cobra.Command{
	Use:   "getPeers",
	Short: "                                 Get the connected peers' information",
	Long:  `Returns the information of connected p2p nodes.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers, err := RPC.GetPeers(context.Background())
		if err != nil {
			fmt.Printf("peers not found: %v\n", err)
			return
		}
		fmt.Printf("Peers: \n%s\n", peers)
	},
}

var getGroupPeersCmd = &cobra.Command{
	Use:   "getGroupPeers",
	Short: "                                 Get all peers' ID within the group",
	Long:  `Returns an ID list of consensus nodes and observer nodes within the specified group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers, err := RPC.GetGroupPeers(context.Background())
		if err != nil {
			fmt.Printf("peers not found: %v\n", err)
			return
		}
		fmt.Printf("Peers: \n%s\n", peers)
	},
}

var getNodeIDListCmd = &cobra.Command{
	Use:   "getNodeIDList",
	Short: "                                 Get ID list of nodes",
	Long:  `Returns an ID list of the node itself and the connected p2p nodes.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers, err := RPC.GetNodeIDList(context.Background())
		if err != nil {
			fmt.Printf("node ID list not found: %v\n", err)
			return
		}
		fmt.Printf("Node ID list: \n%s\n", peers)
	},
}

var getGroupListCmd = &cobra.Command{
	Use:   "getGroupList",
	Short: "                                 Get ID list of groups that the node belongs",
	Long:  `Returns an ID list of groups that the node belongs.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		peers, err := RPC.GetGroupList(context.Background())
		if err != nil {
			fmt.Printf("group IDs list not found: %v\n", err)
			return
		}
		fmt.Printf("Group ID List: \n%s\n", peers)
	},
}

// ========= block access ==========

var getBlockByHashCmd = &cobra.Command{
	Use:   "getBlockByHash",
	Short: "[blockHash]   [true/false]       Query the block by its hash",
	Long: `Returns the block information according to the block hash.
Arguments:
          [blockHash]: hash string
[includeTransactions]: must be "true" or "false".

For example:

    [getBlockByHash] [0x910ea44e2a83618c7cc98456678c9984d94977625e224939b24b3c904794b5ec] [true]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var bhash string
		var includeTx bool

		_, err := isValidHex(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 1 {
			bhash = args[0]
			includeTx = true
		} else {
			bhash = args[0]
			_includeTx, err := strconv.ParseBool(args[1])
			if err != nil {
				fmt.Printf("Arguments error: please check your input: %s%s: %v\n", args[1], info, err)
				return
			}
			includeTx = _includeTx
		}
		peers, err := RPC.GetBlockByHash(context.Background(), bhash, includeTx)
		if err != nil {
			fmt.Printf("block not found: %v\n", err)
			return
		}
		fmt.Printf("Block: \n%s\n", peers)
	},
}

var getBlockByNumberCmd = &cobra.Command{
	Use:   "getBlockByNumber",
	Short: "[blockNumber] [true/false]       Query the block by its number",
	Long: `Returns the block information according to the block number.
Arguments:
       [blockNumber]: can be input in a decimal or in hex(prefix with "0x").
[includeTransactions]: must be "true" or "false".

For example:

    [getBlockByNumber] [0x9] [true]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var bnumber string
		var includeTx bool

		bnum, err := isValidNumber(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 1 {
			bnumber = args[0]
			includeTx = true
		} else {
			bnumber = args[0]
			_includeTx, err := strconv.ParseBool(args[1])
			if err != nil {
				fmt.Printf("Arguments error: please check your input: %s%s: %v\n", args[1], info, err)
				return
			}
			includeTx = _includeTx
		}

		_, err = isOutOfRange(bnum)
		if err != nil {
			fmt.Println(err)
			return
		}

		block, err := RPC.GetBlockByNumber(context.Background(), bnumber, includeTx)
		if err != nil {
			fmt.Printf("block not found: %v\n", err)
			return
		}
		fmt.Printf("Block: \n%s\n", block)
	},
}

var getBlockHashByNumberCmd = &cobra.Command{
	Use:   "getBlockHashByNumber",
	Short: "[blockNumber]                    Query the block hash by its number",
	Long: `Returns the block hash according to the block number.
Arguments:
[blockNumber]: can be input in a decimal format or in hex(prefix with "0x").

For example:

    [getBlockHashByNumber] [0x9]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bnum, err := isValidNumber(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = isOutOfRange(bnum)
		if err != nil {
			fmt.Println(err)
			return
		}

		bnumber := args[0]
		bhash, err := RPC.GetBlockHashByNumber(context.Background(), bnumber)
		if err != nil {
			fmt.Printf("block not found: %v\n", err)
			return
		}
		fmt.Printf("Block Hash: \n%s\n", bhash)
	},
}

// ======== transaction access ========

var getTransactionByHashCmd = &cobra.Command{
	Use:   "getTransactionByHash",
	Short: "[transactionHash]                Query the transaction by its hash",
	Long: `Returns the transaction according to the transaction hash.
Arguments:
[transactionHash]: hash string.

For example:
	
    [getTransactionByHash] [0x7536cf1286b5ce6c110cd4fea5c891467884240c9af366d678eb4191e1c31c6f]
	
For more information please refer:
	
    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := isValidHex(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		txHash := args[0]
		tx, err := RPC.GetTransactionByHash(context.Background(), txHash)
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction: \n%s\n", tx)
	},
}

var getTransactionByBlockHashAndIndexCmd = &cobra.Command{
	Use:   `getTransactionByBlockHashAndIndex`,
	Short: "[blockHash]   [transactionIndex] Query the transaction by block hash and transaction index",
	Long: `Returns transaction information based on block hash and transaction index inside the block.
Arguments:
       [blockHash]: block hash string.
[transactionIndex]: index for the transaction that must be encoded in hex format(prefix with "0x").

For example:

    [getTransactionByBlockHashAndIndex] [0x10bfdc1e97901ed22cc18a126d3ebb8125717c2438f61d84602f997959c631fa] [0x0]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := isValidHex(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		// starts with "0x"
		if !strings.HasPrefix(args[1], "0x") {
			fmt.Println("Arguments error: Not a valid hex string, please check your inpunt: ", args[1], info)
			return
		}

		_, err = isValidNumber(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		bhash := args[0]
		txIndex := args[1]
		tx, err := RPC.GetTransactionByBlockHashAndIndex(context.Background(), bhash, txIndex)
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction: \n%s\n", tx)
	},
}

var getTransactionByBlockNumberAndIndexCmd = &cobra.Command{
	Use:   "getTransactionByBlockNumberAndIndex",
	Short: "[blockNumber] [transactionIndex] Query the transaction by block number and transaction index",
	Long: `Returns transaction information based on block number and transaction index inside the block.
Arguments:
     [blockNumber]: block number encoded in decimal format or in hex(prefix with "0x").
[transactionIndex]: index for the transaction that must be encoded in hex format(prefix with "0x").

For example:

    [getTransactionByBlockNumberAndIndex] [0x9] [0x0]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := isValidNumber(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		// starts with "0x"
		if !strings.HasPrefix(args[1], "0x") {
			fmt.Println("Arguments error: Not a valid hex string, please check your inpunt: ", args[1], info)
			return
		}

		_, err = isValidNumber(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		bnumber := args[0]
		txIndex := args[1]
		tx, err := RPC.GetTransactionByBlockNumberAndIndex(context.Background(), bnumber, txIndex)
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction: \n%s\n", tx)
	},
}

var getTransactionReceiptCmd = &cobra.Command{
	Use:   "getTransactionReceipt",
	Short: "[transactionHash]                Query the transaction receipt by transaction hash",
	Long: `Returns transaction receipt information based on transaction hash.
Arguments:
[transactionHash]: transaction hash string.

For example:

    [getTransactionReceipt] [0x708b5781b62166bd86e543217be6cd954fd815fd192b9a124ee9327580df8f3f]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := isValidHex(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		txHash := args[0]
		tx, err := RPC.GetTransactionReceipt(context.Background(), txHash)
		if err != nil {
			fmt.Printf("transaction receipt not found: %v\n", err)
			return
		}
		fmt.Printf("Transaction Receipt: \n%s\n", tx)
	},
}

var getPendingTransactionsCmd = &cobra.Command{
	Use:   "getPendingTransactions",
	Short: "                                 Get the pending transactions",
	Long:  `Return the transactions that are pending for packaging.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tx, err := RPC.GetPendingTransactions(context.Background())
		if err != nil {
			fmt.Printf("transaction not found: %v\n", err)
			return
		}
		fmt.Printf("Pending Transactions: \n%s\n", tx)
	},
}

var getPendingTxSizeCmd = &cobra.Command{
	Use:   "getPendingTxSize",
	Short: "                                 Get the count of pending transactions",
	Long:  `Return the total count of pending transactions.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tx, err := RPC.GetPendingTxSize(context.Background())
		if err != nil {
			fmt.Printf("transactions not found: %v\n", err)
			return
		}
		fmt.Printf("Pending Transactions Count: \n    hex: %s\n", tx)
		strNum := string(tx)
		bnum, err := toDecimal(strNum[3 : len(strNum)-1])
		if err != nil {
			fmt.Println("The Pending Transactions Count is tot a valid hex string")
			return
		}
		fmt.Println("decimal: ", bnum)
	},
}

// ======== contracts =======

var getCodeCmd = &cobra.Command{
	Use:   "getCode",
	Short: "[contract address]               Get the contract data from contract address",
	Long: `Return contract data queried according to contract address.
Arguments:
[contract address]: contract hash string.

For example:

    [getCode] [0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := isValidHex(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		contractAdd := args[0]
		code, err := RPC.GetCode(context.Background(), contractAdd)
		if err != nil {
			fmt.Printf("This address does not exist: %v\n", err)
			return
		}

		if len(string(code)) < 5 {
			fmt.Println("This address does not exist: ", args[0])
			return
		}

		fmt.Printf("Contract Code: \n%s\n", code)
	},
}

var getTotalTransactionCountCmd = &cobra.Command{
	Use:   "getTotalTransactionCount",
	Short: "                                 Get the totoal transactions and the latest block height",
	Long:  `Returns the current total number of transactions and block height.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		counts, err := RPC.GetTotalTransactionCount(context.Background())
		if err != nil {
			fmt.Printf("information not found: %v\n", err)
			return
		}
		fmt.Printf("Latest Statistics on Transaction and Block Height: \n%s\n", counts)
	},
}

var getSystemConfigByKeyCmd = &cobra.Command{
	Use:   "getSystemConfigByKey",
	Short: "[tx_count_limit/tx_gas_limit]    Get the system configuration through key-value",
	Long: `Returns the system configuration through key-value.
Arguments:
[key to query]: currently only support two key: "tx_count_limit" and "tx_gas_limit".

For example:

    [getSystemConfigByKey] [tx_count_limit]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "tx_count_limit" && args[0] != "tx_gas_limit" {
			fmt.Println("The key not found: ", args[0], ", currently only support [tx_count_limit] and [tx_gas_limit]")
			return
		}
		key := args[0]
		value, err := RPC.GetSystemConfigByKey(context.Background(), key)
		if err != nil {
			fmt.Printf("information not found: %v\n", err)
			return
		}
		fmt.Printf("Result: \n%s\n", value)
	},
}

// ======= contract operation =====

var setSystemConfigByKeyCmd = &cobra.Command{
	Use:   "setSystemConfigByKey",
	Short: "[tx_count_limit/tx_gas_limit]    Set the system configuration through key-value",
	Long: `Returns the system configuration through key-value.
Arguments:
	  [key]: currently only support two key: "tx_count_limit" and "tx_gas_limit".
[key value]: the value of corresponding key.

For example:

    [setSystemConfigByKey] [tx_count_limit] 10000

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html#`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "tx_count_limit" && args[0] != "tx_gas_limit" {
			fmt.Println("The key not found: ", args[0], ", currently only support [tx_count_limit] and [tx_gas_limit]")
			return
		}
		key := args[0]
		value := args[1]
		sysConfig, err := config.NewSystemConfigService(RPC)
		if err != nil {
			fmt.Printf("init systemConfigPrecompiled failed: %v\n", err)
			return
		}
		_, err = sysConfig.SetValueByKey(key, value)
		if err != nil {
			fmt.Printf("SetValueByKey failed: %v\n", err)
			return
		}
		fmt.Printf("Result: \n%s\n", value)
	},
}

// ======= auto completion script =====

var completionCmd = &cobra.Command{
	Use:     "completion [bash|zsh]",
	Short:   "                                 Generate completion script",
	Long: `To load completions:

Bash:

$ source <(./console completion bash)

# To load completions for each session, execute once:
Linux:
  $ ./console completion bash > go_sdk_completion && sudo mv go_sdk_completion /etc/bash_completion.d/
MacOS:
  $ ./console completion bash > /usr/local/etc/bash_completion.d/go_sdk_completion

Zsh:

$ source <(./console completion zsh)

# To load completions for each session, execute once:
$ ./console completion zsh > ~/.go-sdk-completion.sh && echo 'source ~/.go-sdk-completion.sh' >> ~/.zshrc
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			if err != nil {
				fmt.Printf("completion for bash failed: %v\n", err)
				return
			}
		case "zsh":
			err := runCompletionZsh(os.Stdout, cmd.Root())
			if err != nil {
				fmt.Printf("completion for zsh failed: %v\n", err)
				return
			}
		}
	},
}

func init() {
	// add common command
	// TODO: test the bash scripts
	// rootCmd.AddCommand(bashCompletionCmd, zshCompletionCmd)
	// add node command
	rootCmd.AddCommand(getClientVersionCmd, getGroupIDCmd, getBlockNumberCmd, getPbftViewCmd, getSealerListCmd)
	rootCmd.AddCommand(getObserverListCmd, getConsensusStatusCmd, getSyncStatusCmd, getPeersCmd, getGroupPeersCmd)
	rootCmd.AddCommand(getNodeIDListCmd, getGroupListCmd)
	// add block access command
	rootCmd.AddCommand(getBlockByHashCmd, getBlockByNumberCmd, getBlockHashByNumberCmd)
	// add transaction command
	rootCmd.AddCommand(getTransactionByHashCmd, getTransactionByBlockHashAndIndexCmd, getTransactionByBlockNumberAndIndexCmd)
	rootCmd.AddCommand(getTransactionReceiptCmd, getPendingTransactionsCmd, getPendingTxSizeCmd)
	// add contract command
	rootCmd.AddCommand(getCodeCmd, getTotalTransactionCountCmd, getSystemConfigByKeyCmd)
	// add contract command
	rootCmd.AddCommand(setSystemConfigByKeyCmd)
	// add auto completion command
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func toDecimal(hex string) (int, error) {
	i := new(big.Int)
	var flag bool
	i, flag = i.SetString(hex, 16) // octal
	if flag != true {
		return -1, fmt.Errorf("Cannot parse hex string to Int")
	}
	return int(i.Uint64()), nil
}

func isValidNumber(str string) (int, error) {
	var bnum int
	var err error
	// whether contains "." as the float type
	if strings.Contains(str, ".") {
		return -1, fmt.Errorf("Arguments error: please check your inpunt: %s%s", str, info)
	}
	// starts with "0x"
	if strings.HasPrefix(str, "0x") {
		// is hex string
		_, err = strconv.ParseInt(str[2:len(str)], 16, 64)
		if err != nil {
			return -1, fmt.Errorf("Not a valid hex string: arguments error: please check your inpunt: %s%s: %v", str, info, err)
		}
		bnum, err = toDecimal(str[2:len(str)])
		if err != nil {
			return -1, fmt.Errorf("Not a valid hex string: arguments error: please check your inpunt: %s%s", str, info)
		}
	} else {
		bnum, err = strconv.Atoi(str)
		if err != nil {
			return -1, fmt.Errorf("Arguments error: please check your input: %s%s: %v", str, info, err)
		}
	}
	return bnum, nil
}

func isValidHex(str string) (bool, error) {
	// starts with "0x"
	if strings.HasPrefix(str, "0x") {
		if len(str) == 2 {
			return false, fmt.Errorf("Not a valid hex string: arguments error: please check your inpunt: %s%s", str, info)
		}
		// is hex string
		_, err := hexutil.Decode(str)
		if err != nil {
			return false, fmt.Errorf("Not a valid hex string: arguments error: please check your inpunt: %s%s: %v", str, info, err)
		}
		return true, nil
	}
	return false, fmt.Errorf("Arguments error: Not a valid hex string, please check your inpunt: %s%s", str, info)
}

func isOutOfRange(bnum int) (bool, error) {
	// compare with the current block number
	curr, err := RPC.GetBlockNumber(context.Background())
	if err != nil {
		return false, fmt.Errorf("Client error: cannot get the block number: %v", err)
	}
	currStr := string(curr)
	currInt, err := toDecimal(currStr[3 : len(currStr)-1])
	if err != nil {
		return false, fmt.Errorf("Client error: cannot get the block number: %v", err)
	}
	if currInt < bnum {
		return false, fmt.Errorf("BlockNumber does not exist")
	}
	return true, nil
}

func runCompletionZsh(out io.Writer,cmd *cobra.Command) error {
	zshHead := "#compdef kubectl\n"

	_, err := out.Write([]byte(zshHead))
	if err != nil {
		return err
	}

	zshInitialization := `
__kubectl_bash_source() {
	alias shopt=':'
	emulate -L sh
	setopt kshglob noshglob braceexpand

	source "$@"
}

__kubectl_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift

		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__kubectl_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}

__kubectl_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?

	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}

__kubectl_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}

__kubectl_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}

__kubectl_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}

__kubectl_filedir() {
	# Don't need to do anything here.
	# Otherwise we will get trailing space without "compopt -o nospace"
	true
}

autoload -U +X bashcompinit && bashcompinit

# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q 'GNU\|BusyBox'; then
	LWORD='\<'
	RWORD='\>'
fi

__kubectl_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__kubectl_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__kubectl_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__kubectl_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__kubectl_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__kubectl_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__kubectl_type/g" \
	<<'BASH_COMPLETION_EOF'
`
	_, err = out.Write([]byte(zshInitialization))
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = cmd.GenBashCompletion(buf)
	if err != nil {
		return err
	}

	_, err = out.Write(buf.Bytes())
	if err != nil {
		return err
	}

	zshTail := `
BASH_COMPLETION_EOF
}

__kubectl_bash_source <(__kubectl_convert_bash_to_zsh)
`
	_, err = out.Write([]byte(zshTail))
	if err != nil {
		return err
	}

	return nil
}