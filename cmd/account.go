package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/CreatureDev/xrpl-go/client"
	jsonrpcclient "github.com/CreatureDev/xrpl-go/client/jsonrpc"
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

func RunForRequest(req client.XRPLRequest, field string, method string) func(*cobra.Command, []string) {
	return func(*cobra.Command, []string) {
		cl := NewClient()
		rf := reflect.ValueOf(cl).Elem()
		f := rf.FieldByName(field)
		m := f.MethodByName(method)
		results := m.Call([]reflect.Value{reflect.ValueOf(req)})
		if !results[2].IsNil() {
			fmt.Println(results[2].Elem())
		} else {
			printResponse(results[0].Interface())
		}
	}
}

func initRequest(req client.XRPLRequest, field string, method string) {
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
				fmt.Println(req.Method() + ": failed parsing " + tag)
			}

		}
		if !optional {
			newCmd.MarkFlagRequired(tag)
		}
	}

	newCmd.Run = RunForRequest(req, field, method)
}

func NewClient() *client.XRPLClient {
	cfg, err := client.NewJsonRpcConfig(Host)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return jsonrpcclient.NewClient(cfg)
}

func init() {
	rootCmd.AddCommand(accountCmd)

	initRequest(&AccountChannels, "Account", "AccountChannels")
	initRequest(&AccountCurrencies, "Account", "AccountCurrencies")
	initRequest(&AccountInfo, "Account", "AccountInfo")
	initRequest(&AccountLines, "Account", "AccountLines")
	initRequest(&AccountNFTs, "Account", "AccountNFTs")
	initRequest(&AccountObjects, "Account", "AccountObjects")
	initRequest(&AccountOffers, "Account", "AccountOffers")
	initRequest(&AccountTransactions, "Account", "AccountTransactions")
}
