package cmd

import (
	"github.com/CreatureDev/xrpl-go/model/client/clio"
	"github.com/spf13/cobra"
)

var clioCmd = &cobra.Command{
	Use:   "clio",
	Short: "clio commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var ClioServerInfo clio.ServerInfoRequest
var ClioLedger clio.LedgerRequest
var NFTInfo clio.NFTInfoRequest
var NFTHistory clio.NFTHistoryRequest

func init() {
	rootCmd.AddCommand(clioCmd)
	initRequest(&ClioServerInfo, "Clio", "ServerInfo")
	initRequest(&ClioLedger, "Clio", "Ledger")
	initRequest(&NFTInfo, "Clio", "NFTInfo")
	initRequest(&NFTHistory, "Clio", "NFTHistory")
}
