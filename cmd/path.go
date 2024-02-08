package cmd

import (
	"github.com/CreatureDev/xrpl-go/model/client/path"
	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "path and order book commands",
	Run:   func(*cobra.Command, []string) {},
}

var BookOffers path.BookOffersRequest
var DepositAuthorized path.DepositAuthorizedRequest
var NFTBuyOffers path.NFTokenBuyOffersRequest
var NFTSellOffers path.NFTokenSellOffersRequest
var PathFind path.PathFindRequest
var RipplePathFind path.RipplePathFindRequest

func init() {
	rootCmd.AddCommand(pathCmd)

	initRequest(&BookOffers, "Path", "BookOffers")
	initRequest(&DepositAuthorized, "Path", "DepositAuthorized")
	initRequest(&NFTBuyOffers, "Path", "NFTokenBuyOffers")
	initRequest(&NFTSellOffers, "Path", "NFTokenSellOffers")
	initRequest(&PathFind, "Path", "PathFind")
	initRequest(&RipplePathFind, "Path", "RipplePathFind")
}
