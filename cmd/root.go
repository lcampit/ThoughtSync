/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"os"

	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ThoughtSync",
	Short: "A tool to manage your notes",
	Long: `ThoughtSync is a CLI tool that helps command line power users 
  managing their notes. It allows to create and edit notes at 
  the speed of thought from anywhere, without leaving the terminal.`,
}

// Global logger
var Printer = printer.NewPrinter()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	// RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/.config/thoughtsync/config.yaml)")
	RootCmd.PersistentFlags().StringP("vault", "v", config.DEFAULT_VAULT_PATH, "Your notes vault path")
	viper.BindPFlag("vault.path", RootCmd.PersistentFlags().Lookup("vault"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
