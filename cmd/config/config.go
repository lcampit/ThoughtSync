package config

import (
	"os"

	gopath "path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// Configuration file in ~/.config/CONFIG_DIR/CONFIG_FILE
	CONFIG_FILE = "thoughtsync"
	CONFIG_DIR  = "thoughtsync"

	// Keys in the config file with default values
	// Vault path
	VAULT_KEY                     = "vault.path"
	DEFAULT_VAULT_PATH            = "$HOME/thoughtsync-vault"
	VAULT_NOTES_EXTENSION_KEY     = "vault.extension"
	DEFAULT_VAULT_NOTES_EXTENSION = ".md"

	// Journal note format
	JOURNAL_NOTE_FORMAT_KEY = "journal.format"
	DEFAULT_JOURNAL_FORMAT  = "YYYY-MM-DD"

	// Journal directory
	JOURNAL_DIRECTORY_KEY     = "journal.directory"
	DEFAULT_JOURNAL_DIRECTORY = "journal"
)

func InitConfig() {
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)
	viper.SetConfigType("yaml")
	viper.SetConfigName(CONFIG_FILE)
	viper.AddConfigPath(gopath.Join(configDir, CONFIG_DIR))
	viper.AutomaticEnv()

	viper.SetDefault(VAULT_KEY, DEFAULT_VAULT_PATH)
	viper.SetDefault(JOURNAL_NOTE_FORMAT_KEY, DEFAULT_JOURNAL_FORMAT)
	viper.SetDefault(JOURNAL_DIRECTORY_KEY, DEFAULT_JOURNAL_DIRECTORY)
	viper.SetDefault(VAULT_NOTES_EXTENSION_KEY, DEFAULT_VAULT_NOTES_EXTENSION)

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}
}
