/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"fmt"
	gopath "path"
	filepath "path/filepath"

	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/editor"
	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func OpenInboxNote(editor editor.Editor, vaultPath, inboxNotePath, fileExtension string) error {
	completeFilename := gopath.Base(inboxNotePath)

	inboxDirPath := gopath.Dir(inboxNotePath)
	if filepath.Ext(inboxNotePath) != fileExtension {
		completeFilename += fileExtension
	}
	folderPath := gopath.Join(vaultPath, inboxDirPath)
	err := path.EnsurePresent(folderPath, completeFilename)
	if err != nil {
		return fmt.Errorf("error in ensure present for dir %s, file %s: %v", folderPath, completeFilename, err)
	}
	fullPath := gopath.Join(folderPath, completeFilename)
	err = editor.Edit(fullPath)
	if err != nil {
		return fmt.Errorf("error in write: %v", err)
	}
	return nil
}

func init() {
	editor := editor.NewEditor()
	InboxCmd := &cobra.Command{
		Use:   "inbox",
		Short: "Creates and opens your inbox note in your $EDITOR",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			vaultPath := viper.GetString(config.VAULT_KEY)
			fileExtension := viper.GetString(config.VAULT_NOTES_EXTENSION_KEY)
			inboxNotePath := viper.GetString(config.INBOX_NOTE_KEY)
			err := OpenInboxNote(editor, vaultPath, inboxNotePath, fileExtension)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	RootCmd.AddCommand(InboxCmd)
}
