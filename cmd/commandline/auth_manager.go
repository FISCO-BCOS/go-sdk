package commandline

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/precompiled/permission"
	"github.com/spf13/cobra"
)

var getCommitteeInfo = &cobra.Command{
	Use:   "getCommitteeInfo",
	Short: "                                   get Committee info",
	Long: `get Committee info.
For example:

    [getCommitteeInfo]

For more information please refer:

  	https://fisco-bcos-doc.readthedocs.io/zh_CN/latest/docs/develop/sdk/java_sdk/rpc_api.html#getcommitteeinfo`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		permissionService, err := permission.NewPermissionService(RPC)
		if err != nil {
			fmt.Printf("listPermissionManager failed, permission.NewPermissionService err:%v\n", err)
			return
		}
		managers, err := permissionService.ListPermissionManager()
		if err != nil {
			fmt.Printf("listPermissionManager failed, permissionService.ListPermissionManager err: %v\n", err)
			return
		}
		jsonStr, err := ListToJSONStr(managers, "managers")
		if err != nil {
			fmt.Printf("listPermissionManager failed, ListToJsonStr err: %v\n", err)
			return
		}
		fmt.Println(jsonStr)
	},
}
