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

	gopath "path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func OpenTodayNote(editor editor.Editor, vaultPath string) error {
	journalPath := gopath.Join(vaultPath, "journal")
	filename := fmt.Sprintf("%s.md", date.Today())
	err := path.EnsurePresent(journalPath, filename)
	if err != nil {
		return fmt.Errorf("failed to ensure present: %w", err)
	}
	filePath := gopath.Join(journalPath, filename)
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
			return OpenTodayNote(editor, vaultPath)
		},
	}
	RootCmd.AddCommand(todayCmd)
}
