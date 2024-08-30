/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/lcampit/ThoughtSync/cmd/config"
	"github.com/lcampit/ThoughtSync/cmd/repository"

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
		Printer.Info("vault worktree clean, nothing to sync")
		return nil
	}

	err = repository.AddAllAndCommit(commitMessage)
	if err != nil {
		return err
	}
	if !remoteEnabled {
		Printer.Info("remote option is not enabled, skipping push")
		return nil
	}

	if !skipPush {
		err = repository.Push()
		return err
	} else {
		Printer.Info("skipping push")
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
	if status != "" {
		Printer.Info(fmt.Sprintf("%s\n", status))
	} else {
		Printer.Info("vault is clean")
	}
	return nil
}

// VaultGitPush pushses changes to the vault remote git repo
func VaultGitPush(repository repository.Repository) error {
	err := repository.Push()
	return err
}

// VaultGitPull pulls changes from the vault remote
// into the local repository
func VaultGitPull(repository repository.Repository) error {
	err := repository.Pull()
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
		Run: func(cmd *cobra.Command, args []string) {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				Printer.CustomError("git sync is not enabled in your config file")
				return
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			commitMessage := viper.GetString(config.GIT_COMMIT_MESSAGE_KEY)
			remoteEnabled := viper.GetBool(config.GIT_REMOTE_ENABLED_KEY)
			skipPush, _ := cmd.Flags().GetBool("no-push")
			useSSHAuth := viper.GetBool(config.GIT_AUTH_SSH_KEY)
			remoteName := viper.GetString(config.GIT_REMOTE_NAME_KEY)

			repository, err := repository.OpenRepository(vaultPath, remoteName, useSSHAuth)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			err = SyncWithGit(repository, commitMessage, remoteEnabled, skipPush)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Print out the git status of your vault repo",
		Run: func(cmd *cobra.Command, args []string) {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				Printer.CustomError("git sync is not enabled in your config file")
				return
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			repository, err := repository.OpenRepository(vaultPath, "", false)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			err = VaultGitStatus(repository)
			if err != nil {
				Printer.PlainError(err)
				return
			}
		},
	}
	syncCmd.Flags().Bool("no-push", viper.GetBool(config.GIT_REMOTE_ENABLED_KEY), "do not perform push after git commit")

	pushCmd := &cobra.Command{
		Use:   "push",
		Short: "Push changes to the vault remote git repo",
		Run: func(cmd *cobra.Command, args []string) {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				Printer.CustomError("git sync is not enabled in your config file")
				return
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			useSSHAuth := viper.GetBool(config.GIT_AUTH_SSH_KEY)
			remoteName := viper.GetString(config.GIT_REMOTE_NAME_KEY)

			repository, err := repository.OpenRepository(vaultPath, remoteName, useSSHAuth)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			err = VaultGitPush(repository)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull changes from the vault remote git repo",
		Run: func(cmd *cobra.Command, args []string) {
			gitSyncEnabled := viper.GetBool(config.GIT_SYNC_ENABLED_KEY)
			if !gitSyncEnabled {
				Printer.CustomError("git sync is not enabled in your config file")
				return
			}
			vaultPath := viper.GetString(config.VAULT_KEY)
			useSSHAuth := viper.GetBool(config.GIT_AUTH_SSH_KEY)
			remoteName := viper.GetString(config.GIT_REMOTE_NAME_KEY)

			repository, err := repository.OpenRepository(vaultPath, remoteName, useSSHAuth)
			if err != nil {
				Printer.PlainError(err)
				return
			}
			err = VaultGitPull(repository)
			if err != nil {
				Printer.PlainError(err)
			}
		},
	}

	gitCmd.AddCommand(syncCmd, statusCmd, pushCmd, pullCmd)
	RootCmd.AddCommand(gitCmd)
}
