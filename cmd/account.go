package cmd

import (
	"fmt"

	"github.com/CreatureDev/xrpl-go/client"
	"github.com/CreatureDev/xrpl-go/model/client/account"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Account string

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "fetch account info",
}

func RunInfo(cl *client.XRPLClient, command *cobra.Command, args []string) {
	var req *account.AccountInfoRequest
	req = &account.AccountInfoRequest{
		Account: types.Address(Account),
	}

	res, _, _ := cl.Account.AccountInfo(req)

	fmt.Printf("%+v\n", res)
}

func init() {
	infoCmd.Run = withClient(RunInfo)
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().StringVarP(&Account, "account", "a", "", "account to fetch info of")
	infoCmd.MarkFlagRequired("account")
	viper.BindPFlag("account", infoCmd.PersistentFlags().Lookup("account"))
}
