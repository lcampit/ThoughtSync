package cmd

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/printer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CONFIG_KEY_ORDER = []string{"vault", "journal", "git"}

func getConfig(printer printer.Printer) {
	configKeys := config.GetAllConfigKeys()
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

func setConfig(printer printer.Printer, configKey, configValue string) {
	if !slices.Contains(config.GetAllConfigKeys(), configKey) {
		printer.Error(fmt.Sprintf("%s is not a valid configuration key", configKey))
	}

	config.SetConfig(configKey, configValue)
}

func init() {
	printer := printer.NewPrinter()
	getConfigCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"c"},
		Run: func(cmd *cobra.Command, args []string) {
			getConfig(printer)
		},
	}
	setConfigCmd := &cobra.Command{
		Use:   "set <config-key> <config-value>",
		Short: "Sets a configuration option to a value",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			setConfig(printer, args[0], args[1])
		},
	}
	getConfigCmd.AddCommand(setConfigCmd)
	RootCmd.AddCommand(getConfigCmd)
}
