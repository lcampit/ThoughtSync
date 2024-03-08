package date

import (
	"ThoughtSync/cmd/date"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DateTestSuite struct {
	suite.Suite
}

func (suite *DateTestSuite) SetupSuite() {
}

func (suite *DateTestSuite) TestTodayDateFormat() {
	date := date.Today()
	// Asserts date format is YYYY-MM-DD
	assert.Len(suite.T(), date, 10)
}

func (suite *DateTestSuite) TearDownSuite() {
}

func TestDateTestSuite(t *testing.T) {
	suite.Run(t, new(DateTestSuite))
}
