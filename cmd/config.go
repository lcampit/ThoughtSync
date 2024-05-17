package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lcampit/ThoughtSync/cmd/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CONFIG_KEY_ORDER = []string{"vault", "journal", "git"}

func getConfig(printer printer.Printer) {
	configKeys := viper.AllKeys()
	groupedConfigKeys := make(map[string][]string)
	for _, key := range configKeys {
		splitKey := strings.Split(key, ".")
		keyGroup := splitKey[0]
		subKey := splitKey[1]
		groupedConfigKeys[keyGroup] = append(groupedConfigKeys[keyGroup], fmt.Sprintf("%s: %s", subKey, viper.GetString(key)))
		sort.Strings(groupedConfigKeys[keyGroup])
	}

	for _, key := range CONFIG_KEY_ORDER {
		printer.ConfigMainKey(key)
		for _, subkey := range groupedConfigKeys[key] {
			subKeyValue := fmt.Sprintf("  %s", subkey)
			printer.ConfigSubKey(subKeyValue)
		}
	}
}

func init() {
	printer := printer.NewPrinter()
	getConfigCmd := &cobra.Command{
		Use: "config",
		Run: func(cmd *cobra.Command, args []string) {
			getConfig(printer)
		},
	}
	RootCmd.AddCommand(getConfigCmd)
}
