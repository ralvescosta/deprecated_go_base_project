package errors

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InternalErrorTestSuite struct {
	suite.Suite
}

func TestInternalErrorTestSuite(t *testing.T) {
	suite.Run(t, new(InternalErrorTestSuite))
}

func (s *InternalErrorTestSuite) TestNewInternalError() {
	err := NewInternalError("some error")
	s.Error(err)
	s.IsType(InternalError{}, err)
}

func (s *InternalErrorTestSuite) TestNewInternalErrorError() {
	err := NewInternalError("some error")
	s.Equal("some error", err.Error())
}
