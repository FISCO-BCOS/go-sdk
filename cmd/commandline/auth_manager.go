package commandline

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/precompiled/auth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var getDeployAuthType = &cobra.Command{
	Use:   "getDeployAuthType",
	Short: "                                   get global deploy auth type",
	Long: `get global deploy auth type.
For example:

    [getDeployAuthType]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#getdeployauthtype`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("getDeployAuthType failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.GetDeployAuthType()
		if err != nil {
			fmt.Printf("getDeployAuthType failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var checkDeployAuth = &cobra.Command{
	Use:   "checkDeployAuth",
	Short: "[accountAddress]                   check the account whether this account can deploy contract.",
	Long: `check the account whether this account can deploy contract.
Arguments:
         [checkDeployAuth]: string

For example:

    [checkDeployAuth] [0x112fb844934c794a9e425dd6b4e57eff1b519f17]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#checkdeployauth`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accountAddress := args[0]
		if !IsValidAccount(accountAddress) {
			fmt.Printf("the format of accountAddress %v is unvalid\n", accountAddress)
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("checkDeployAuth failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.CheckDeployAuth(common.HexToAddress(accountAddress))
		if err != nil {
			fmt.Printf("checkDeployAuth failed,  err: %v\n", err)
			return
		}
		fmt.Println(*result)
	},
}

var checkMethodAuth = &cobra.Command{
	Use:   "checkMethodAuth",
	Short: "[contractAddr][funcSelector][account]  check the contract interface func whether this account can call.",
	Long: `check the contract interface func whether this account can call.
Arguments:
         [contractAddr]: the contractAddress
         [funcSelector]: interface func selector of contract, 4 bytes
         [account]: the account to check

For example:

    [checkMethodAuth] [0x112fb844934c794a9e425dd6b4e57eff1b519f17][c53057b4][0x112fb844934c794a9e425dd6b4e57eff1b519f17]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#checkmethodauth`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		contractAddr := args[0]
		if !IsValidAccount(contractAddr) {
			fmt.Printf("the format of contractAddr %v is unvalid\n", contractAddr)
			return
		}

		funcSelectorStr := args[1]
		funcByte, err := hex.DecodeString(funcSelectorStr)
		if err != nil {
			fmt.Printf("the format of funcSelector %v is unvalid , hex decode err\n", funcSelectorStr)
			return
		}
		if len(funcByte) != 4 {
			fmt.Printf("the format of funcSelector %v is unvalid\n", funcSelectorStr)
			return
		}

		var funcs [4]byte
		funcs[0] = funcByte[0]
		funcs[1] = funcByte[1]
		funcs[2] = funcByte[2]
		funcs[3] = funcByte[3]

		accountAddress := args[2]
		if !IsValidAccount(accountAddress) {
			fmt.Printf("the format of accountAddress %v is unvalid\n", accountAddress)
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("checkMethodAuth failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.CheckMethodAuth(common.HexToAddress(contractAddr), funcs, common.HexToAddress(accountAddress))
		if err != nil {
			fmt.Printf("checkMethodAuth failed,  err: %v\n", err)
			return
		}
		fmt.Println(*result)
	},
}

var getAdmin = &cobra.Command{
	Use:   "getAdmin",
	Short: "[contractAddr]                     get a specific contract admin.",
	Long: `get a specific contract admin.
Arguments:
         [contractAddr]: string

For example:

    [getAdmin] [0x112fb844934c794a9e425dd6b4e57eff1b519f17]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#getadmin`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contractAddr := args[0]
		if !IsValidAccount(contractAddr) {
			fmt.Printf("the format of contractAddr %v is unvalid\n", contractAddr)
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("getAdmin failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.GetAdmin(common.HexToAddress(contractAddr))
		if err != nil {
			fmt.Printf("getAdmin failed,  err: %v\n", err)
			return
		}
		fmt.Println(*result)
	},
}

var resetAdmin = &cobra.Command{
	Use:   "resetAdmin",
	Short: "[newAdmin][contractAddr]           submit a proposal of resetting contract admin, only governor can call it.",
	Long: `submit a proposal of resetting contract admin, only governor can call it.
Arguments:
         [newAdmin]: new admin address
         [contractAddr]: the address of contract which will propose to reset admin

For example:

    [resetAdmin] [0x112fb844934c794a9e425dd6b4e57eff1b519f17][0x112fb844934c794a9e425dd6b4e57eff1b519f17]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#resetadmin`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		newAdmin := args[0]
		if !IsValidAccount(newAdmin) {
			fmt.Printf("the format of newAdmin %v is unvalid\n", newAdmin)
			return
		}
		contractAddr := args[1]
		if !IsValidAccount(contractAddr) {
			fmt.Printf("the format of contractAddr %v is unvalid\n", contractAddr)
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("resetAdmin failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.ResetAdmin(common.HexToAddress(newAdmin), common.HexToAddress(contractAddr))
		if err != nil {
			fmt.Printf("resetAdmin failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var updateGovernor = &cobra.Command{
	Use:   "updateGovernor",
	Short: "[accountAddress][weight]           apply for update governor, only governor can call it.",
	Long: `apply for update governor, only governor can call it.
Arguments:
         [accountAddress]: string
         [weight]: uint32

For example:

    [updateGovernor] [0x112fb844934c794a9e425dd6b4e57eff1b519f17] [1]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#updategovernor`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		accountAddress := args[0]
		if !IsValidAccount(accountAddress) {
			fmt.Printf("the format of accountAddress %v is unvalid\n", accountAddress)
			return
		}

		weight, err := strconv.ParseInt(args[1], 0, 0)
		if err != nil {
			fmt.Println("weight should be integer")
			return
		}
		if weight < 0 || weight > math.MaxInt32 {
			fmt.Println("weight failed, must be >= 0")
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("updateGovernor failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.UpdateGovernor(common.HexToAddress(accountAddress), uint32(weight))
		if err != nil {
			fmt.Printf("updateGovernor failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var setRate = &cobra.Command{
	Use:   "setRate",
	Short: "[participatesRate][winRate]        apply set participate rate and win rate. only governor can call it.",
	Long: `apply set participate rate and win rate. only governor can call it.
Arguments:
         [participatesRate]: uint8
         [winRate]: uint8

For example:

    [setRate] [1] [1]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#setrate`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		participatesRateStr, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("participatesRate should be integer")
			return
		}
		if participatesRateStr < 0 || participatesRateStr > 100 {
			fmt.Println("participatesRate failed, must be must be in the range of [0, 100]")
			return
		}
		participatesRate := uint8(participatesRateStr)

		winRate, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("winRate should be integer")
			return
		}
		if winRate < 0 || winRate > 100 {
			fmt.Println("winRate failed, must be must be in the range of [0, 100]")
			return
		}
		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("setRate failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.SetRate(participatesRate, uint8(winRate))
		if err != nil {
			fmt.Printf("setRate failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var setDeployAuthType = &cobra.Command{
	Use:   "setDeployAuthType",
	Short: "[deployAuthType]                   submit a proposal of setting deploy contract auth type, only governor can call it.",
	Long: `submit a proposal of setting deploy contract auth type, only governor can call it.
Arguments:
         [deployAuthType]: uint8 1-whitelist; 2-blacklist

For example:

    [setDeployAuthType] [1]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#setdeployauthtype`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		num, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("deployAuthType should be integer")
			return
		}
		if num != 1 && num != 2 {
			fmt.Println("deployAuthType failed, must be [1-whitelist or 2-blacklist] ")
			return
		}
		deployAuthType := uint8(num)

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("setDeployAuthType failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.SetDeployAuthType(deployAuthType)
		if err != nil {
			fmt.Printf("setDeployAuthType failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var modifyDeployAuth = &cobra.Command{
	Use:   "modifyDeployAuth",
	Short: "[accountAddress][openFlag]         submit a proposal of adding deploy contract auth for account, only governor can call it.",
	Long: `submit a proposal of adding deploy contract auth for account, only governor can call it.
Arguments:
         [accountAddress]: account address string
         [openFlag]: true-open; false-close

For example:

    [modifyDeployAuth] [0x112fb844934c794a9e425dd6b4e57eff1b519f17][true]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#modifydeployauth`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		accountAddress := args[0]
		if !IsValidAccount(accountAddress) {
			fmt.Printf("the format of accountAddress %v is unvalid\n", accountAddress)
			return
		}

		openFlag, err := strconv.ParseBool(args[1])
		if err != nil {
			fmt.Printf("openFlag failed, must be [true-open; false-close] %s", args[1])
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("modifyDeployAuth failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.ModifyDeployAuth(common.HexToAddress(accountAddress), openFlag)
		if err != nil {
			fmt.Printf("modifyDeployAuth failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var revokeProposal = &cobra.Command{
	Use:   "revokeProposal",
	Short: "[proposalId]                       revoke proposal, only governor can call it.",
	Long: `revoke proposal, only governor can call it.
Arguments:
         [proposalId]: proposal id

For example:

    [revokeProposal] [1]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#revokeproposal`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		proposalId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("proposalId should be integer")
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("revokeProposal failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.RevokeProposal(*big.NewInt(proposalId))
		if err != nil {
			fmt.Printf("revokeProposal failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var voteProposal = &cobra.Command{
	Use:   "voteProposal",
	Short: "[proposalId][agree]                unified vote, only governor can call it.",
	Long: `unified vote, only governor can call it.
Arguments:
         [proposalId]: proposal id
         [agree]: true or false

For example:

    [voteProposal] [1]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#voteproposal`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		proposalId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("proposalId should be integer")
			return
		}

		agree, err := strconv.ParseBool(args[1])
		if err != nil {
			fmt.Printf("agree failed, must be [true or false] %s", args[1])
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("voteProposal failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.VoteProposal(*big.NewInt(proposalId), agree)
		if err != nil {
			fmt.Printf("voteProposal failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var setMethodAuthType = &cobra.Command{
	Use:   "setMethodAuthType",
	Short: "[contractAddr][funcSelector][authType]                set a specific contract's method auth type, only contract admin can call it.",
	Long: `set a specific contract's method auth type, only contract admin can call it.
Arguments:
         [contractAddr]: the contract address to set auth
         [funcSelector]:  interface func selector of contract, 4 bytes
         [authType]: uint8 1-whitelist; 2-blacklist

For example:

    [setMethodAuthType] [0x][0x9cc3ca0f][1]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#setmethodauthtype`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		accountAddress := args[0]
		if !IsValidAccount(accountAddress) {
			fmt.Printf("the format of accountAddress %v is unvalid\n", accountAddress)
			return
		}

		funcSelectorStr := args[1]
		funcByte, err := hex.DecodeString(funcSelectorStr)
		if err != nil {
			fmt.Printf("the format of funcSelector %v is unvalid , hex decode err\n", funcSelectorStr)
			return
		}
		if len(funcByte) != 4 {
			fmt.Printf("the format of funcSelector %v is unvalid\n", funcSelectorStr)
			return
		}

		var funcs [4]byte
		funcs[0] = funcByte[0]
		funcs[1] = funcByte[1]
		funcs[2] = funcByte[2]
		funcs[3] = funcByte[3]

		num, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("deployAuthType should be integer")
			return
		}
		if num != 1 && num != 2 {
			fmt.Println("deployAuthType failed, must be [1-whitelist or 2-blacklist] ")
			return
		}
		authType := uint8(num)

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("setMethodAuthType failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.SetMethodAuthType(common.HexToAddress(accountAddress), funcs, authType)
		if err != nil {
			fmt.Printf("setMethodAuthType failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

var setMethodAuth = &cobra.Command{
	Use:   "setMethodAuth",
	Short: "[contractAddr][funcSelector][account][isOpen]         set a specific contract's method ACL, only contract admin can call it",
	Long: `set a specific contract's method ACL, only contract admin can call it.
Arguments:
         [contractAddr]: the contract address to set auth
         [funcSelector]:  interface func selector of contract, 4 bytes
         [account]: account
         [isOpen]: true , false

For example:

    [setMethodAuth] [0x][0x9cc3ca0f][0x][false]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#setmethodauth`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		contractAddr := args[0]
		if !IsValidAccount(contractAddr) {
			fmt.Printf("the format of contractAddr %v is unvalid\n", contractAddr)
			return
		}

		funcSelectorStr := args[1]
		funcByte, err := hex.DecodeString(funcSelectorStr)
		if err != nil {
			fmt.Printf("the format of funcSelector %v is unvalid , hex decode err\n", funcSelectorStr)
			return
		}
		if len(funcByte) != 4 {
			fmt.Printf("the format of funcSelector %v is unvalid\n", funcSelectorStr)
			return
		}

		var funcs [4]byte
		funcs[0] = funcByte[0]
		funcs[1] = funcByte[1]
		funcs[2] = funcByte[2]
		funcs[3] = funcByte[3]

		accountAddr := args[2]
		if !IsValidAccount(accountAddr) {
			fmt.Printf("the format of accountAddr %v is unvalid\n", accountAddr)
			return
		}

		isOpen, err := strconv.ParseBool(args[3])
		if err != nil {
			fmt.Printf("isOpen failed, must be [true or false] %s", args[1])
			return
		}

		authManagerService, err := auth.NewAuthManagerService(RPC)
		if err != nil {
			fmt.Printf("setMethodAuth failed,  err:%v\n", err)
			return
		}
		result, err := authManagerService.SetMethodAuth(common.HexToAddress(contractAddr), funcs, common.HexToAddress(accountAddr), isOpen)
		if err != nil {
			fmt.Printf("setMethodAuth failed,  err: %v\n", err)
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand( /*getCommitteeInfo, getProposalInfo, */ getDeployAuthType, checkDeployAuth, checkMethodAuth, getAdmin)
	rootCmd.AddCommand(updateGovernor, setRate, setDeployAuthType, modifyDeployAuth, resetAdmin, revokeProposal, voteProposal)
	rootCmd.AddCommand(setMethodAuthType, setMethodAuth)
}
