/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/config"
	"ThoughtSync/cmd/date"
	"ThoughtSync/cmd/editor"
	"ThoughtSync/cmd/path"
	"fmt"
	"time"

	gopath "path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func OpenTodayNote(editor editor.Editor, vaultPath, filename string) error {
	journalPath := gopath.Join(vaultPath, "journal")
	filenameMd := fmt.Sprintf("%s.md", filename)
	err := path.EnsurePresent(journalPath, filenameMd)
	if err != nil {
		return fmt.Errorf("failed to ensure present: %w", err)
	}
	filePath := gopath.Join(journalPath, filenameMd)
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
			format := viper.GetString(config.JOURNAL_NOTE_FORMAT)
			filename, err := date.Format(time.Now(), format)
			if err != nil {
				return fmt.Errorf("error getting journal filename: %w", err)
			}
			return OpenTodayNote(editor, vaultPath, filename)
		},
	}
	RootCmd.AddCommand(todayCmd)
}
