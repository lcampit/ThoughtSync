/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cat(filePath string) error {
	return nil
}

func init() {
	readCmd := &cobra.Command{
		Use:   "read",
		Short: "Quickly select and print out to the terminal a note content",
		Run: func(cmd *cobra.Command, args []string) {
			vaultPath := viper.GetString(config.VAULT_KEY)
			files, err := path.ListAllContents(vaultPath)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			var selectedFile path.FileInfo
			selected, err := fuzzyfinder.Find(files, func(i int) string {
				// selectedFile = files[i]
				// return files[i].Path
				digit := strconv.Itoa(i)
				return digit
			}, fuzzyfinder.WithHeader("Select a note to read"),
				fuzzyfinder.WithPreviewWindow(func(i int, width, height int) string {
					return fmt.Sprintf("Preview: %s", files[i].Name)
				}))
			if err == fuzzyfinder.ErrAbort {
				// Nothing selected
				return
			}
			if err != nil {
				Printer.PlainError(err)
				return
			}

			fmt.Print(selected)
			fmt.Print(selectedFile)

			// err := cat(selected)
			// if err != nil {
			// 	Printer.PlainError(err)
			// }
		},
	}

	RootCmd.AddCommand(readCmd)
}
