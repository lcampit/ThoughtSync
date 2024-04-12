package cmd

import (
	"errors"
	"testing"

	"github.com/lcampit/ThoughtSync/cmd"
	"github.com/lcampit/ThoughtSync/mocks/cmd/repository"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NewGitCmdTestSuite struct {
	CmdTestSuite
	repository *repository.Repository
}

func (suite *NewGitCmdTestSuite) SetupSuite() {
	suite.repository = repository.NewRepository(suite.T())
}

func (suite *NewGitCmdTestSuite) SetupTest() {
	suite.repository = repository.NewRepository(suite.T())
}

func (suite *NewGitCmdTestSuite) TestGitStatusCmdReturns() {
	suite.repository.On("GetStatusAsString").Return("status", nil)
	err := cmd.VaultGitStatus(suite.repository)

	suite.Assert().Nil(err)
	suite.repository.AssertCalled(suite.T(), "GetStatusAsString")
}

func (suite *NewGitCmdTestSuite) TestGitStatusCmdReturnsError() {
	suite.repository.On("GetStatusAsString").Return("", errors.New("failed"))
	err := cmd.VaultGitStatus(suite.repository)

	suite.Assert().NotNil(err)
	suite.repository.AssertCalled(suite.T(), "GetStatusAsString")
}

func (suite *NewGitCmdTestSuite) TestGitPushCmdReturns() {
	suite.repository.On("Push").Return(nil)
	err := cmd.VaultGitPush(suite.repository)

	suite.Assert().Nil(err)
	suite.repository.AssertCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitPushCmdReturnsError() {
	suite.repository.On("Push").Return(errors.New("failed"))
	err := cmd.VaultGitPush(suite.repository)

	suite.Assert().NotNil(err)
	suite.repository.AssertCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitPullCmdReturns() {
	suite.repository.On("Pull").Return(nil)
	err := cmd.VaultGitPull(suite.repository)

	suite.Assert().Nil(err)
	suite.repository.AssertCalled(suite.T(), "Pull")
}

func (suite *NewGitCmdTestSuite) TestGitPullCmdReturnsError() {
	suite.repository.On("Pull").Return(errors.New("failed"))
	err := cmd.VaultGitPull(suite.repository)

	suite.Assert().NotNil(err)
	suite.repository.AssertCalled(suite.T(), "Pull")
}

func (suite *NewGitCmdTestSuite) TestGitSyncCmd() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(nil)
	suite.repository.On("IsClean").Return(false, nil)
	suite.repository.On("AddAllAndCommit", mock.Anything).Return(nil)
	suite.repository.On("Push").Return(nil)

	err := cmd.SyncWithGit(suite.repository, commitMessage, true, false)

	suite.Assert().Nil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertCalled(suite.T(), "IsClean")
	suite.repository.AssertCalled(suite.T(), "AddAllAndCommit", commitMessage)
	suite.repository.AssertCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitSyncFailsOnPullError() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(errors.New("failed"))

	err := cmd.SyncWithGit(suite.repository, commitMessage, true, false)

	suite.Assert().NotNil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertNotCalled(suite.T(), "IsClean")
	suite.repository.AssertNotCalled(suite.T(), "AddAllAndCommit", mock.Anything)
	suite.repository.AssertNotCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitSyncFailsOnIsCleanError() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(nil)
	suite.repository.On("IsClean").Return(false, errors.New("failed"))

	err := cmd.SyncWithGit(suite.repository, commitMessage, true, false)

	suite.Assert().NotNil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertCalled(suite.T(), "IsClean")
	suite.repository.AssertNotCalled(suite.T(), "AddAllAndCommit", mock.Anything)
	suite.repository.AssertNotCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitSyncFailsOnAddAllAndCommitError() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(nil)
	suite.repository.On("IsClean").Return(false, nil)
	suite.repository.On("AddAllAndCommit", commitMessage).Return(errors.New("failed"))

	err := cmd.SyncWithGit(suite.repository, commitMessage, true, false)

	suite.Assert().NotNil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertCalled(suite.T(), "IsClean")
	suite.repository.AssertCalled(suite.T(), "AddAllAndCommit", commitMessage)
	suite.repository.AssertNotCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitSyncFailsOnPushError() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(nil)
	suite.repository.On("IsClean").Return(false, nil)
	suite.repository.On("AddAllAndCommit", commitMessage).Return(nil)
	suite.repository.On("Push").Return(errors.New("failed"))

	err := cmd.SyncWithGit(suite.repository, commitMessage, true, false)

	suite.Assert().NotNil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertCalled(suite.T(), "IsClean")
	suite.repository.AssertCalled(suite.T(), "AddAllAndCommit", commitMessage)
	suite.repository.AssertCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitSyncCmdDoesNotPushOnRemoteNotEnabled() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(nil)
	suite.repository.On("IsClean").Return(false, nil)
	suite.repository.On("AddAllAndCommit", mock.Anything).Return(nil)

	err := cmd.SyncWithGit(suite.repository, commitMessage, false, false)

	suite.Assert().Nil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertCalled(suite.T(), "IsClean")
	suite.repository.AssertCalled(suite.T(), "AddAllAndCommit", commitMessage)
	suite.repository.AssertNotCalled(suite.T(), "Push")
}

func (suite *NewGitCmdTestSuite) TestGitSyncCmdDoesNotPushOnSkipPush() {
	commitMessage := "message"
	suite.repository.On("Pull").Return(nil)
	suite.repository.On("IsClean").Return(false, nil)
	suite.repository.On("AddAllAndCommit", mock.Anything).Return(nil)

	err := cmd.SyncWithGit(suite.repository, commitMessage, true, true)

	suite.Assert().Nil(err)

	suite.repository.AssertCalled(suite.T(), "Pull")
	suite.repository.AssertCalled(suite.T(), "IsClean")
	suite.repository.AssertCalled(suite.T(), "AddAllAndCommit", commitMessage)
	suite.repository.AssertNotCalled(suite.T(), "Push")
}

func TestGitCmdTestSuite(t *testing.T) {
	suite.Run(t, new(NewGitCmdTestSuite))
}
