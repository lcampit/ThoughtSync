/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/config"
	"ThoughtSync/cmd/repository"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SyncWithGit adds all files to staging, commits with a given
// message and pushes to remote if the pushToRemote flag is true
func SyncWithGit(repository repository.Repository, commitMessage string, remoteEnabled, skipPush bool) error {
	err := repository.Pull()
	if err != nil {
		return err
	}

	isClean, err := repository.IsClean()
	if err != nil {
		return err
	}
	if isClean {
		fmt.Println("vault worktree clean, nothing to sync")
		return nil
	}

	err = repository.AddAllAndCommit(commitMessage)
	if err != nil {
		return err
	}
	if !remoteEnabled {
		fmt.Println("remote option is not enabled, skipping push")
		return nil
	}

	if !skipPush {
		err = repository.Push()
		return err
	} else {
		fmt.Println("skipping push")
	}

	return nil
}

// VaultGitStatus prints to stdout the status
// of the vault git repo, i.e. its current working tree
func VaultGitStatus(repository repository.Repository) error {
	status, err := repository.GetStatusAsString()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", status)
	return nil
}

// VaultGitPush pushses changes to the vault remote git repo
func VaultGitPush(repository repository.Repository) error {
	err := repository.Push()
	return err
}

func init() {
	gitCmd := &cobra.Command{
		Use:   "git",
		Short: "Git related commands",
	}
	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Stage, commit and push all changes in your note vault to your remote repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				return fmt.Errorf("git sync is not enabled in your config file")
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			commitMessage := viper.GetString(config.GIT_COMMIT_MESSAGE_KEY)
			remoteEnabled := viper.GetBool(config.GIT_REMOTE_ENABLED_KEY)
			skipPush, _ := cmd.Flags().GetBool("no-push")
			useSSHAuth := viper.GetBool(config.GIT_AUTH_SSH_KEY)
			remoteName := viper.GetString(config.GIT_REMOTE_NAME_KEY)

			repository, err := repository.OpenRepository(vaultPath, remoteName, useSSHAuth)
			if err != nil {
				return err
			}
			return SyncWithGit(repository, commitMessage, remoteEnabled, skipPush)
		},
	}

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Print out the git status of your vault repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				return fmt.Errorf("git sync is not enabled in your config file")
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			repository, err := repository.OpenRepository(vaultPath, "", false)
			if err != nil {
				return err
			}
			return VaultGitStatus(repository)
		},
	}
	syncCmd.Flags().Bool("no-push", viper.GetBool(config.GIT_REMOTE_ENABLED_KEY), "do not perform push after git commit")

	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Push changes to the vault remote git repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				return fmt.Errorf("git sync is not enabled in your config file")
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			useSSHAuth := viper.GetBool(config.GIT_AUTH_SSH_KEY)
			remoteName := viper.GetString(config.GIT_REMOTE_NAME_KEY)

			repository, err := repository.OpenRepository(vaultPath, remoteName, useSSHAuth)
			if err != nil {
				return err
			}
			return VaultGitPush(repository)
		},
	}
	gitCmd.AddCommand(syncCmd, statusCmd, pushCmd)
	RootCmd.AddCommand(gitCmd)
}
