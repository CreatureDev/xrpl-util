package cmd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/CreatureDev/xrpl-go/client"
	"github.com/CreatureDev/xrpl-go/model/client/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var specifier string

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var AccountChannels account.AccountChannelsRequest
var AccountCurrencies account.AccountCurrenciesRequest
var AccountInfo account.AccountInfoRequest
var AccountLines account.AccountLinesRequest
var AccountNFTs account.AccountNFTsRequest
var AccountObjects account.AccountObjectsRequest
var AccountOffers account.AccountOffersRequest
var AccountTransactions account.AccountTransactionsRequest

func printResponse(res any) {
	s, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(s))
}

func RunForRequest(req client.XRPLRequest) func(cl *client.XRPLClient, command *cobra.Command, args []string) {
	method := req.Method()
	switch method {
	case "account_channels":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			acr := req.(*account.AccountChannelsRequest)
			res, _, err := cl.Account.AccountChannels(acr)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_currencies":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			acr := req.(*account.AccountCurrenciesRequest)
			res, _, err := cl.Account.AccountCurrencies(acr)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_info":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			air := req.(*account.AccountInfoRequest)
			res, _, err := cl.Account.AccountInfo(air)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_lines":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			alr := req.(*account.AccountLinesRequest)
			res, _, err := cl.Account.AccountLines(alr)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_nfts":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			anr := req.(*account.AccountNFTsRequest)
			res, _, err := cl.Account.AccountNFTs(anr)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_objects":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			aor := req.(*account.AccountObjectsRequest)
			res, _, err := cl.Account.AccountObjects(aor)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_offers":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			aor := req.(*account.AccountOffersRequest)
			res, _, err := cl.Account.AccountOffers(aor)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				printResponse(res)
			}
		}
	case "account_tx":
		return func(cl *client.XRPLClient, command *cobra.Command, args []string) {
			atr := req.(*account.AccountTransactionsRequest)
			res, _, err := cl.Account.AccountTransactions(atr)
			if err != nil {
				fmt.Println("tx " + err.Error())
			} else {
				printResponse(res)
			}
		}
	}
	return func(*client.XRPLClient, *cobra.Command, []string) {
		fmt.Println(method + " unimplemented")
	}
}

func initRequest(req client.XRPLRequest) {
	var newCmd = &cobra.Command{
		Use: req.Method(),
	}

	accountCmd.AddCommand(newCmd)

	rv := reflect.ValueOf(req).Elem()
	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i)
		t := rv.Type().Field(i)
		tag := t.Tag.Get("json")
		optional := strings.Contains(tag, "omitempty")
		if tag == "" {
			tag = strings.ToLower(t.Name)
		} else if strings.Contains(tag, ",") {
			tag = tag[:strings.IndexByte(tag, ',')]
		}
		if !v.CanAddr() {
			fmt.Println("CANNOT ADDR " + t.Tag.Get("json"))
		}
		switch v.Kind() {
		case reflect.Bool:
			newCmd.Flags().BoolVar((*bool)(v.Addr().UnsafePointer()), tag, false, tag)
			viper.BindPFlag(tag, newCmd.Flags().Lookup(tag))

		case reflect.Uint:
			newCmd.Flags().UintVar((*uint)(v.Addr().UnsafePointer()), tag, 0, tag)
			viper.BindPFlag(tag, newCmd.Flags().Lookup(tag))

		case reflect.Int:
			newCmd.Flags().IntVar((*int)(v.Addr().UnsafePointer()), tag, 0, tag)
			viper.BindPFlag(tag, newCmd.Flags().Lookup(tag))

		case reflect.String:
			newCmd.Flags().StringVar((*string)(v.Addr().UnsafePointer()), tag, "", tag)
			viper.BindPFlag(tag, newCmd.Flags().Lookup(tag))

		case reflect.Interface:
			switch tag {
			case "marker":
			//byte string
			case "ledger_index":
				newCmd.Flags().StringVar(&specifier, tag, "", tag)
				viper.BindPFlag(tag, newCmd.Flags().Lookup(tag))
			default:
				fmt.Println("failed parsing " + tag)
			}

		}
		if !optional {
			newCmd.MarkFlagRequired(tag)
		}
	}
	newCmd.Run = withClient(RunForRequest(req))

}

func init() {
	rootCmd.AddCommand(accountCmd)

	initRequest(&AccountChannels)
	initRequest(&AccountCurrencies)
	initRequest(&AccountInfo)
	initRequest(&AccountLines)
	initRequest(&AccountNFTs)
	initRequest(&AccountObjects)
	initRequest(&AccountOffers)
	initRequest(&AccountTransactions)
}
