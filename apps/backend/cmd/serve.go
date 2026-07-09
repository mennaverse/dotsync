package cmd

import (
	"fmt"

	"github.com/mennaverse/dotsync/apps/backend/consts"
	"github.com/mennaverse/dotsync/apps/backend/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the backend server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bindAddress, _ := cmd.Flags().GetString("bind")

		err := web.StartServer(bindAddress)
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	},
}

func init() {
	serveCmd.Flags().StringP("bind", "b", consts.BindAddressDefault, "Bind address for the server")
	serveCmd.Flags().String("server-name", consts.ServerNameDefault, "Server name for the server")

	viper.BindPFlag("bind", serveCmd.Flags().Lookup("bind"))
	viper.BindPFlag("server-name", serveCmd.Flags().Lookup("server-name"))

	rootCmd.AddCommand(serveCmd)
}
