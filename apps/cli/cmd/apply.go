package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply [profile]",
	Short: "Apply a specific dotfile configuration by specifying its name.",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]

		fmt.Printf("Applying profile: %s\n", profile)
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
