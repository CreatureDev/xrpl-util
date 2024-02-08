package cmd

import (
	"github.com/CreatureDev/xrpl-go/model/client/ledger"
	"github.com/spf13/cobra"
)

var ledgerCmd = &cobra.Command{
	Use:   "ledger",
	Short: "ledger commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var Ledger ledger.LedgerRequest
var LedgerClosed ledger.LedgerClosedRequest
var LedgerCurrent ledger.LedgerCurrentRequest
var LedgerData ledger.LedgerDataRequest
var LedgerEntry ledger.LedgerEntryRequest

func init() {
	rootCmd.AddCommand(ledgerCmd)

	initRequest(&Ledger, "Ledger", "Ledger")
	initRequest(&LedgerClosed, "Ledger", "LedgerClosed")
	initRequest(&LedgerCurrent, "Ledger", "LedgerCurrent")
	initRequest(&LedgerData, "Ledger", "LedgerData")
	initRequest(&LedgerEntry, "Ledger", "LedgerEntry")
}
