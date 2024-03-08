package cmd

import (
	"ThoughtSync/mocks/cmd/editor"
	"fmt"
	"os"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CmdTestSuite struct {
	suite.Suite
	editor    *editor.Editor
	vaultPath string
}

func (suite *CmdTestSuite) SetupSuite() {
	tmpFolder, err := os.MkdirTemp("", "vault-test-folder")
	if err != nil {
		fmt.Printf("error in creating tmp directory %s", err)
		suite.T().FailNow()
	}
	suite.vaultPath = tmpFolder
	suite.editor = editor.NewEditor(suite.T())
	suite.editor.On("Edit", mock.Anything).Return(nil)
}

func (suite *CmdTestSuite) TearDownSuite() {
	os.RemoveAll(suite.vaultPath)
}
