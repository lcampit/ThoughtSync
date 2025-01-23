/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/editor"
	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func OpenVaultDir(editor editor.Editor, vaultPath string) error {
	err := path.CreateFolders(vaultPath)
	if err != nil {
		return fmt.Errorf("error creating dirs to vault path: %s: %v", vaultPath, err)
	}
	err = editor.Edit(vaultPath)
	if err != nil {
		return fmt.Errorf("error in write: %v", err)
	}
	return nil
}

func init() {
	editor := editor.NewEditor()
	OpenCmd := &cobra.Command{
		Use:     "vault",
		Aliases: []string{"v"},
		Short:   "Opens the vault directory in your $EDITOR",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			vaultPath := viper.GetString(config.VAULT_KEY)
			err := OpenVaultDir(editor, vaultPath)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	RootCmd.AddCommand(OpenCmd)
}
