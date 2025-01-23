/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	readCmd := &cobra.Command{
		Use:   "read",
		Short: "Quickly select and print out to the terminal a note content",
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

			f, err := os.Open(files[indexSelected[0]].Path)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			defer f.Close()
			_, err = io.Copy(os.Stdout, f)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	RootCmd.AddCommand(readCmd)
}
