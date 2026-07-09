package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Opens the register page in the default web browser for user registration.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opening the register page in the default web browser...")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
