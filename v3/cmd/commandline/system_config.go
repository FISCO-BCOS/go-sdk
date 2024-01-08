package commandline

import (
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/v3/precompiled/config"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/spf13/cobra"
)

var configSet = mapset.NewSet[string]("tx_count_limit", "tx_gas_limit", "consensus_leader_period", "rpbft_epoch_block_num", "auth_check_status", "compatibility_version", "feature_rpbft_epoch_sealer_num", "feature_rpbft_epoch_block_num", "feature_rpbft", "feature_balance_precompiled", "feature_rpbft_notify_rotate")

var setSystemConfigByKey = &cobra.Command{
	Use:   "setSystemConfigByKey",
	Short: "[system_configuration_item]        Set the system configuration through key-value",
	Long: fmt.Sprintf(`Returns the system configuration through key-value.
Arguments:
	  [key]: currently only support four key: %s.
[key value]: the value of corresponding key.

For example:
    setSystemConfigByKey tx_count_limit 10000`, configSet.String()),
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if configSet.Contains(args[0]) {
			fmt.Printf("The key not found: %s, currently only support %v", args[0], configSet)
			return
		}
		key := args[0]
		value := args[1]
		sysConfig, err := config.NewSystemConfigService(RPC)
		if err != nil {
			fmt.Printf("setSystemConfigByKey failed, config.NewSystemConfigService err: %v\n", err)
			return
		}
		result, err := sysConfig.SetValueByKey(key, value)
		if err != nil {
			fmt.Printf("setSystemConfigByKey failed, sysConfig.SetValueByKey err: %v\n", err)
			return
		}
		if result != 0 {
			fmt.Printf("setSystemConfigByKey failed, the result is: %v", result)
			return
		}
		fmt.Println("success")
	},
}

func init() {
	rootCmd.AddCommand(setSystemConfigByKey)
}
