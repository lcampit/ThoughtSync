package path

import (
	"fmt"
	"os"
	gopath "path"
	"testing"

	"github.com/lcampit/ThoughtSync/cmd/path"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PathTestSuite struct {
	suite.Suite
	TmpFolderPath string
	TestFolders   string
	TestFile      string
}

func (suite *PathTestSuite) SetupSuite() {
	tmpFolder, err := os.MkdirTemp("", "vault-test-folder")
	if err != nil {
		fmt.Printf("error in creating tmp directory %s", err)
		suite.T().FailNow()
	}
	suite.TmpFolderPath = tmpFolder
	suite.TestFolders = "vault-folder1/vault-folder2"
	suite.TestFile = "test-file"
}

func (suite *PathTestSuite) TestIsExistsPositive() {
	assert.True(suite.T(), path.IsExist(suite.TmpFolderPath))
}

func (suite *PathTestSuite) TestIsExistsNegative() {
	assert.False(suite.T(), path.IsExist(""))
}

func (suite *PathTestSuite) TestCreateFolders() {
	err := path.CreateFolders(gopath.Join(suite.TmpFolderPath, suite.TestFolders))
	assert.Nil(suite.T(), err)
}

func (suite *PathTestSuite) TestCreateFile() {
	filePath := gopath.Join(suite.TmpFolderPath, suite.TestFile)
	err := path.CreateFile(filePath)
	assert.Nil(suite.T(), err)
}

func (suite *PathTestSuite) TestEnsureExits() {
	tmpPath := gopath.Join(suite.TmpFolderPath, suite.TestFolders)
	err := path.EnsurePresent(tmpPath, suite.TestFile)
	assert.Nil(suite.T(), err)
	fullPath := gopath.Join(tmpPath, suite.TestFile)
	assert.True(suite.T(), path.IsExist(fullPath))
}

func TestPathTestSuite(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}

func (suite *PathTestSuite) TearDownSuite() {
	os.RemoveAll(suite.TmpFolderPath)
}
