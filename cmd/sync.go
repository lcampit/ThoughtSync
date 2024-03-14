/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/config"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SyncWithGit adds all files to staging, commits with a given
// message and pushes to remote if the pushToRemote flag is true
func SyncWithGit(vaultPath, commitMessage string, pushToRemote bool) error {
	cmd := exec.Command("git", "add", "-A")
	cmd.Dir = vaultPath
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-a", "-m", commitMessage)
	cmd.Dir = vaultPath
	_, err = cmd.Output()
	if err != nil {
		return err
	}

	if pushToRemote {
		cmd = exec.Command("git", "push")
		cmd.Dir = vaultPath
		_, err = cmd.Output()
	}
	return err
}

func init() {
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Stage, commit and push all changes in your note vault to your remote repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitSyncEnabled := viper.GetBool("git.enabled")
			if !gitSyncEnabled {
				return fmt.Errorf("git sync is not enabled in your config file")
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			commitMessage := viper.GetString(config.GIT_COMMIT_MESSAGE_KEY)
			pushToRemote := viper.GetBool(config.GIT_REMOTE_ENABLED_KEY)
			return SyncWithGit(vaultPath, commitMessage, pushToRemote)
		},
	}
	RootCmd.AddCommand(syncCmd)
}
