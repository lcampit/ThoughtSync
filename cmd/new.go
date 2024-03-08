/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/editor"
	"ThoughtSync/cmd/path"
	"fmt"

	"github.com/spf13/cobra"
)

func NewNote(editor editor.Editor, vaultPath, noteType, filename string) error {
	folderPath := fmt.Sprintf("%s/%s", vaultPath, noteType)
	err := path.EnsurePresent(folderPath, filename)
	if err != nil {
		return fmt.Errorf("error in ensure present for dir %s, file %s: %v", folderPath, filename, err)
	}
	fullPath := fmt.Sprintf("%s/%s", folderPath, filename)
	err = editor.Edit(fullPath)
	if err != nil {
		return fmt.Errorf("error in write: %v", err)
	}
	return nil
}

// newCmd represents the new command

func init() {
	editor := editor.NewEditor()

	newCmd := &cobra.Command{
		Use:   "new -t <type> <filename>",
		Short: "Creates and opens the given file in your $EDITOR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]
			vaultPath, _ := cmd.Flags().GetString("vault")
			noteType, _ := cmd.Flags().GetString("type")
			return NewNote(editor, vaultPath, noteType, filename)
		},
	}
	newCmd.Flags().StringP("type", "t", "", "Folder of the note vault to put the new note in")

	RootCmd.AddCommand(newCmd)
}
