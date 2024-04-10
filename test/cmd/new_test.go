package cmd

import (
	"ThoughtSync/cmd"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NewNoteTestSuite struct {
	CmdTestSuite
}

func (suite *NewNoteTestSuite) TestNewNoteCmd() {
	err := cmd.NewNote(suite.editor, suite.vaultPath, "", "test", ".txt")

	filenameWithExtension := "test.txt"

	suite.editor.AssertCalled(suite.T(), "Edit",
		mock.MatchedBy(func(expected string) bool { return strings.Contains(expected, filenameWithExtension) }))
	suite.Assert().Nil(err)
}

func TestNewNoteTestSuite(t *testing.T) {
	suite.Run(t, new(NewNoteTestSuite))
}
