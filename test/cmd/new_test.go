package cmd

import (
	"ThoughtSync/cmd"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NewNoteTestSuite struct {
	CmdTestSuite
}

func (suite *NewNoteTestSuite) TestNewNoteCmd() {
	err := cmd.NewNote(suite.editor, suite.vaultPath, "", "test.txt")
	suite.Assert().Nil(err)
}

func TestNewNoteTestSuite(t *testing.T) {
	suite.Run(t, new(NewNoteTestSuite))
}
