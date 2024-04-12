package repository

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Repository interface {
	AddAllAndCommit(commitMessage string) error
	Push() error
	Pull() error
	IsClean() (bool, error)
	GetStatusAsString() (string, error)
}

type repository struct {
	repo       *git.Repository
	remoteName string
	useSSH     bool
}

// OpenRepository opens an existing
// git reposotory with the given options
func OpenRepository(path, remoteName string, useSSH bool) (Repository, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return &repository{repo: repo, remoteName: remoteName, useSSH: useSSH}, nil
}

// IsClean returns true if the repository
// work tree is clean, false otherwise
func (r *repository) IsClean() (bool, error) {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return false, err
	}

	status, err := worktree.Status()
	if err != nil {
		return false, err
	}

	return status.IsClean(), nil
}

// Pull pulls changes from the remote
// into the local git repository
func (r *repository) Pull() error {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return err
	}

	pullOptions := &git.PullOptions{
		RemoteName: r.remoteName,
	}

	if r.useSSH {
		authMethod, err := ssh.DefaultAuthBuilder("git")
		if err != nil {
			return err
		}

		pullOptions.Auth = authMethod
	}

	err = worktree.Pull(pullOptions)
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}

func (r *repository) AddAllAndCommit(commitMessage string) error {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return err
	}
	_, err = worktree.Add(".")
	if err != nil {
		return err
	}
	commitOptions := &git.CommitOptions{
		All: true,
	}
	_, err = worktree.Commit(commitMessage, commitOptions)
	if err != nil {
		return err
	}
	return nil
}

// Push pushes local changes into the
// remote git repository
func (r *repository) Push() error {
	pushOptions := &git.PushOptions{
		RemoteName: r.remoteName,
	}

	if r.useSSH {
		authMethod, err := ssh.DefaultAuthBuilder("git")
		if err != nil {
			return err
		}

		pushOptions.Auth = authMethod
	}
	err := r.repo.Push(pushOptions)
	return err
}

// GetStatusAsString returns the work tree status
// as a string, one per file, stating their changes
// returns an empty string on clean work tree
func (r *repository) GetStatusAsString() (string, error) {
	worktree, err := r.repo.Worktree()
	if err != nil {
		return "", err
	}
	status, err := worktree.Status()
	if err != nil {
		return "", err
	}
	return status.String(), nil
}
