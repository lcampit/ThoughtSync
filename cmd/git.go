/*
Copyright © 2024 Leonardo Campitelli leonardo932.campitelli@gmail.com
*/
package cmd

import (
	"ThoughtSync/cmd/config"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
)

// SyncWithGit adds all files to staging, commits with a given
// message and pushes to remote if the pushToRemote flag is true
func SyncWithGit(vaultPath, commitMessage string, pushToRemote bool) error {
	repo, err := git.PlainOpen(vaultPath)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	status, err := worktree.Status()
	if err != nil {
		return err
	}

	if status.IsClean() {
		fmt.Printf("vault worktree clean, nothing to sync")
		return nil
	}

	_, err = worktree.Add(".")
	if err != nil {
		return err
	}

	_, err = worktree.Commit(viper.GetString(config.GIT_COMMIT_MESSAGE_KEY), &git.CommitOptions{All: true})
	if err != nil {
		return err
	}

	if pushToRemote {
		err = repo.Push(&git.PushOptions{})
		return err
	}

	return nil
}

func VaultGitStatus(vaultPath string) error {
	repo, err := git.PlainOpen(vaultPath)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	status, err := worktree.Status()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", status)
	return nil
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
			pushToRemote := viper.GetBool(config.GIT_REMOTE_ENABLED_KEY)
			return SyncWithGit(vaultPath, commitMessage, pushToRemote)
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
			return VaultGitStatus(vaultPath)
		},
	}
	syncCmd.Flags().BoolP("push", "p", viper.GetBool(config.GIT_REMOTE_ENABLED_KEY), "Folder of the note vault to put the new note in")
	viper.BindPFlag(config.GIT_REMOTE_ENABLED_KEY, syncCmd.Flags().Lookup("push"))

	gitCmd.AddCommand(syncCmd, statusCmd)
	RootCmd.AddCommand(gitCmd)
}
