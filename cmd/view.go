package cmd

import (
	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/tree"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetVaultView(vaultPath string) error {
	return tree.GetTree(vaultPath)
}

func init() {
	viewCmd := &cobra.Command{
		Use:   "view",
		Short: "Print out the vault contents",
		RunE: func(cmd *cobra.Command, args []string) error {
			vaultPath := viper.GetString(config.VAULT_KEY)
			return GetVaultView(vaultPath)
		},
	}

	RootCmd.AddCommand(viewCmd)
}
