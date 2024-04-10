package cmd

import (
	"ThoughtSync/cmd"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TodayNoteTestSuite struct {
	CmdTestSuite
}

func (suite *TodayNoteTestSuite) TestNewNoteCmd() {
	err := cmd.OpenTodayNote(suite.editor, suite.vaultPath, "2006-02-01", ".md")
	suite.Assert().Nil(err)
	filenameWithExtension := "2006-02-01.md"

	suite.editor.AssertCalled(suite.T(), "Edit",
		mock.MatchedBy(func(expected string) bool { return strings.Contains(expected, filenameWithExtension) }))
}

func TestTodayNoteTestSuite(t *testing.T) {
	suite.Run(t, new(TodayNoteTestSuite))
}
