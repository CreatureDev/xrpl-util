package cmd

import (
	"fmt"
	"os"

	"github.com/CreatureDev/xrpl-go/client"
	jsonrpcclient "github.com/CreatureDev/xrpl-go/client/jsonrpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Host string

var rootCmd = &cobra.Command{
	Use:   "xrpl-util [flags] [command]",
	Short: "xrpl command line utility",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Host, "host", "", "https://s.altnet.rippletest.net:51234/", "host xrpl server to connect to")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

type clientCmd func(cl *client.XRPLClient, cmd *cobra.Command, args []string)
type runCmd func(cmd *cobra.Command, args []string)

func withClient(c clientCmd) runCmd {
	return func(cmd *cobra.Command, args []string) {
		cfg, err := client.NewJsonRpcConfig(Host)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		cl := jsonrpcclient.NewClient(cfg)
		c(cl, cmd, args)
	}
}
