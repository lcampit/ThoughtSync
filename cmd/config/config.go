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
	VAULT_KEY          = "vault.path"
	DEFAULT_VAULT_PATH = "$HOME/thoughtsync-vault"

	// Journal note format
	JOURNAL_NOTE_FORMAT_KEY = "journal.format"
	DEFAULT_JOURNAL_FORMAT  = "YYYY-MM-DD"

	// Journal directory
	JOURNAL_DIRECTORY_KEY     = "journal.directory"
	DEFAULT_JOURNAL_DIRECTORY = "journal"

	// Git syncing
	GIT_SYNC_ENABLED_KEY       = "git.enable"
	DEFAULT_GIT_SYNC_ENABLED   = false
	GIT_COMMIT_MESSAGE_KEY     = "git.commit-message"
	DEFAULT_GIT_COMMIT_MESSAGE = "thoughtsync: Synced with git"
	GIT_REMOTE_ENABLED_KEY     = "git.remote"
	DEFAULT_GIT_REMOTE_ENABLED = false
	GIT_AUTH_SSH_KEY           = "git.ssh"
	DEFAULT_GIT_AUTH_SSH       = false
	GIT_REMOTE_NAME_KEY        = "git.remote-name"
	DEFAULT_GIT_REMOTE_NAME    = "origin"
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
	viper.SetDefault(GIT_SYNC_ENABLED_KEY, DEFAULT_GIT_SYNC_ENABLED)
	viper.SetDefault(GIT_COMMIT_MESSAGE_KEY, DEFAULT_GIT_COMMIT_MESSAGE)
	viper.SetDefault(GIT_REMOTE_ENABLED_KEY, DEFAULT_GIT_REMOTE_ENABLED)
	viper.SetDefault(GIT_AUTH_SSH_KEY, DEFAULT_GIT_AUTH_SSH)
	viper.SetDefault(GIT_REMOTE_NAME_KEY, DEFAULT_GIT_REMOTE_NAME)

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}
}
