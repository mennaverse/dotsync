package cmd

import (
	"fmt"

	"github.com/mennaverse/dotsync/apps/backend/web"
	"github.com/spf13/cobra"
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
	serveCmd.Flags().StringP("bind", "b", "localhost:8067", "Bind address for the server")

	rootCmd.AddCommand(serveCmd)
}
