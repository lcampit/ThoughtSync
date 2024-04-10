/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"fmt"
	gopath "path"
	"time"

	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/date"
	"github.com/lcampit/ThoughtSync/cmd/editor"
	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// OpenTodayNote opens the today note with name filename in the vault directory
// vaultJournalPath using the editor provider
func OpenTodayNote(editor editor.Editor, vaultJournalPath, filename, extension string) error {
	filenameMd := filename + extension
	err := path.EnsurePresent(vaultJournalPath, filenameMd)
	if err != nil {
		return fmt.Errorf("failed to ensure present: %w", err)
	}
	filePath := gopath.Join(vaultJournalPath, filenameMd)
	err = editor.Edit(filePath)
	if err != nil {
		return fmt.Errorf("error in editing file: %w", err)
	}
	return nil
}

func init() {
	editor := editor.NewEditor()
	todayCmd := &cobra.Command{
		Use:   "today",
		Short: "Quickly edit the journal note for today",
		RunE: func(cmd *cobra.Command, args []string) error {
			vaultPath := viper.GetString(config.VAULT_KEY)
			format := viper.GetString(config.JOURNAL_NOTE_FORMAT_KEY)
			journalDir := viper.GetString(config.JOURNAL_DIRECTORY_KEY)
			vaultJournalPath := gopath.Join(vaultPath, journalDir)
			filename, err := date.Format(time.Now(), format)
			fileExtension := viper.GetString(config.VAULT_NOTES_EXTENSION_KEY)
			if err != nil {
				return fmt.Errorf("error getting journal filename: %w", err)
			}
			return OpenTodayNote(editor, vaultJournalPath, filename, fileExtension)
		},
	}
	RootCmd.AddCommand(todayCmd)
}
