package cmd

import (
	"github.com/CreatureDev/xrpl-go/model/client/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server info commands",
	Run:   func(*cobra.Command, []string) {},
}

var Fee server.FeeRequest
var Manifest server.ManifestRequest
var ServerInfo server.ServerInfoRequest
var ServerState server.ServerStateRequest

func init() {
	rootCmd.AddCommand(serverCmd)

	initRequest(&Fee, "Server", "Fee")
	initRequest(&Manifest, "Server", "Manifest")
	initRequest(&ServerInfo, "Server", "ServerInfo")
	initRequest(&ServerState, "Server", "ServerState")
}
