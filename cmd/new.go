/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/config"
	"ThoughtSync/cmd/editor"
	"ThoughtSync/cmd/path"
	"fmt"

	gopath "path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewNote opens the given file filename in vaultType, optionally in directory
// noteType, using the editor provided
func NewNote(editor editor.Editor, vaultPath, noteType, filename string) error {
	folderPath := gopath.Join(vaultPath, noteType)
	err := path.EnsurePresent(folderPath, filename)
	if err != nil {
		return fmt.Errorf("error in ensure present for dir %s, file %s: %v", folderPath, filename, err)
	}
	fullPath := gopath.Join(folderPath, filename)
	err = editor.Edit(fullPath)
	if err != nil {
		return fmt.Errorf("error in write: %v", err)
	}
	return nil
}

func init() {
	editor := editor.NewEditor()
	newCmd := &cobra.Command{
		Use:   "new -t <type> <filename>",
		Short: "Creates and opens the given file in your $EDITOR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]
			vaultPath := viper.GetString(config.VAULT_KEY)
			noteType, _ := cmd.Flags().GetString("type")
			return NewNote(editor, vaultPath, noteType, filename)
		},
	}
	newCmd.Flags().StringP("type", "t", "", "Folder of the note vault to put the new note in")

	RootCmd.AddCommand(newCmd)
}
