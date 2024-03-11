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
	DEFAULT_VAULT_PATH = "thoughtsync-vault"
)

func InitConfig() {
	configDir, err := os.UserConfigDir()
	cobra.CheckErr(err)
	viper.SetConfigType("yaml")
	viper.SetConfigName(CONFIG_FILE)
	viper.AddConfigPath(gopath.Join(configDir, CONFIG_DIR))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}
}
