package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dsy",
	Short: "Dotsync CLI is a command-line interface for managing dotfiles and configurations.",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("dsy")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("DSY")
	viper.AutomaticEnv()

	rootCmd.Flags().StringP("repo", "r", "", "Dotsync Repository URL for dotfiles")
	viper.BindPFlag("repo", rootCmd.Flags().Lookup("repo"))

	if err := viper.ReadInConfig(); err != nil {
		// Handle error reading config file
		fmt.Printf("Error reading config file: %v\n", err)
	}
}
