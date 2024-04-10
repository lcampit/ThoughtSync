package date

import (
	"testing"
	"time"

	"github.com/lcampit/ThoughtSync/cmd/date"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DateTestSuite struct {
	suite.Suite
}

func (suite *DateTestSuite) SetupSuite() {
}

func (suite *DateTestSuite) TestDateFormatDefaultFormat() {
	date, err := date.Format(time.Now(), "YYYY-MM-DD")
	// Asserts date format is YYYY-MM-DD
	assert.Len(suite.T(), date, 10)
	assert.Nil(suite.T(), err)
}

func (suite *DateTestSuite) TearDownSuite() {
}

func TestDateTestSuite(t *testing.T) {
	suite.Run(t, new(DateTestSuite))
}
