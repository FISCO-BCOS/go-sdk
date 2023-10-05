package commandline

import (
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/precompiled/crud"
	"github.com/spf13/cobra"
)

var createTable = &cobra.Command{
	Use:   "createTable",
	Short: "[tableName] [keyName] [valueName]  create table",
	Long: `createTable.
Arguments:
          [tableName]: string
          [keyName]: string
          [valueName]: string

For example:

    [createTable] [table1] [key1] [value1]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/console.html#addobserver`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		tableName := args[0]
		key := args[1]
		value := args[2]
		var valueFields = []string{}
		valueFields = append(valueFields, value)
		crudService, err := crud.NewCRUDService(RPC)
		if err != nil {
			fmt.Printf("set failed, crud.NewCRUDService err: %v\n", err)
			return
		}
		result, err := crudService.CreateTable(tableName, key, valueFields)
		if err != nil {
			fmt.Printf("set failed, crudService.CreateTable err: %v\n", err)
			return
		}
		if result != 0 {
			fmt.Println("create table failed")
			fmt.Println("result:", result)
			return
		}
		fmt.Println(DefaultSuccessMessage)
	},
}

var set = &cobra.Command{
	Use:   "set",
	Short: "[tableName] [keyName] [valueName]  set key value",
	Long: `set key value.
Arguments:
          [tableName]: string
          [keyName]: string
          [valueName]: string

For example:

    [set] [table1] [key1] [value1]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/console.html#addsealer`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		tableName := args[0]
		key := args[1]
		value := args[2]
		var valueFields = []string{}
		valueFields = append(valueFields, value)
		entry := crud.Entry{
			Key:    key,
			Fields: valueFields,
		}
		crudService, err := crud.NewCRUDService(RPC)
		if err != nil {
			fmt.Printf("set failed, crud.NewCRUDService err: %v\n", err)
			return
		}
		result, err := crudService.Insert(tableName, entry)
		if err != nil {
			fmt.Printf("set failed, crudService.Insert err: %v\n", err)
			return
		}
		if result != 1 {
			fmt.Println("set failed:", result)
			return
		}
		fmt.Println(DefaultSuccessMessage)
	},
}

var get = &cobra.Command{
	Use:   "get",
	Short: "[tableName] [keyName]              get key value",
	Long: `get key value.
Arguments:
          [tableName]: string
          [keyName]: string

For example:

    [get] [tableName] [key1]

For more information please refer:

    https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/console.html#removenode`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		tableName := args[0]
		key := args[1]
		crudService, err := crud.NewCRUDService(RPC)
		if err != nil {
			fmt.Printf("get failed, consensus.NewConsensusService err: %v\n", err)
			return
		}
		result, err := crudService.Select0(tableName, key)
		if err != nil {
			fmt.Printf("get failed, consensusService.AddObserver err: %v\n", err)
			return
		}
		if len(result.Fields) == 0 {
			fmt.Println("get failed")
			return
		}
		for _, value := range result.Fields {
			fmt.Println("get key value:", value)
		}
		fmt.Println(DefaultSuccessMessage)
	},
}

func init() {
	rootCmd.AddCommand(createTable, set, get)
}
