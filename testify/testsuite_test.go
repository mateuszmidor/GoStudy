package testify_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// define test suite
type exampleTestSuite struct {
	suite.Suite

	// test suite internal state
	startAtFive *int
}

// executed once before running entire suite
func (s *exampleTestSuite) SetupSuite() {
	// Log() is printed on test failure/ when running verbose mode with: go test -v .
	s.T().Log("SetupSuite")
	s.startAtFive = new(int)
}

// executed before running every test
func (s *exampleTestSuite) SetupTest() {
	s.T().Log("SetupTest")
	*s.startAtFive = 5
}

// executed after running every test
func (s *exampleTestSuite) TearDownTest() {
	s.T().Log("TearDownTest")
	*s.startAtFive = 0
}

// executed once after running entire suite
func (s *exampleTestSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
	s.startAtFive = nil
}

// test1
func (s *exampleTestSuite) TestMemoryAllocated() {
	s.NotNil(s.startAtFive)
}

// test2
func (s exampleTestSuite) TestStateProperlySet() {
	s.Equal(5, *s.startAtFive)
}

// run entire test suite
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(exampleTestSuite))
}
