package cmd

import (
	"ThoughtSync/cmd"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TodayNoteTestSuite struct {
	CmdTestSuite
}

func (suite *TodayNoteTestSuite) TestNewNoteCmd() {
	err := cmd.OpenTodayNote(suite.editor, suite.vaultPath, "2006-02-01")
	suite.Assert().Nil(err)
}

func TestTodayNoteTestSuite(t *testing.T) {
	suite.Run(t, new(TodayNoteTestSuite))
}
