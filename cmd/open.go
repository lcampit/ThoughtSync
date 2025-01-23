/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/editor"
	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	editor := editor.NewEditor()
	OpenCmd := &cobra.Command{
		Use:     "open",
		Aliases: []string{"o"},
		Short:   "Opens a note in your $EDITOR",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			vaultPath := viper.GetString(config.VAULT_KEY)
			files, err := path.ListAllFiles(vaultPath)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			filenames := make([]string, 0)
			for _, file := range files {
				pathWithoutPrefix, _ := strings.CutPrefix(file.Path, vaultPath+"/")
				filenames = append(filenames, pathWithoutPrefix)
			}

			finder, err := fzf.New(fzf.WithLimit(1))
			if err != nil {
				Printer.PlainError(err)
				return
			}

			indexSelected, err := finder.Find(filenames,
				func(i int) string {
					return filenames[i]
				})

			if err == fzf.ErrAbort {
				// Nothing selected
				return
			}
			if err != nil {
				Printer.PlainError(err)
				return
			}

			err = editor.Edit(files[indexSelected[0]].Path)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	RootCmd.AddCommand(OpenCmd)
}
