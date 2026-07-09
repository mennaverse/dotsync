package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dsyback",
	Short: "Dotsync Backend is a command-line interface for the server managing dotfiles and configurations.",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetConfigName("dsyback")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("DSYBACK")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// Handle error reading config file
		fmt.Printf("Error reading config file: %v\n", err)
	}
}
