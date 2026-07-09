package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [program]",
	Short: "Install dotfiles and programs from the repository",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		programId := args[0]

		fmt.Printf("Installing program: %s\n", programId)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
