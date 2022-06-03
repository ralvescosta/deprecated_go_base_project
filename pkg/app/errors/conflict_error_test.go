package errors

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConflictErrTestSuite struct {
	suite.Suite
}

func TestConclictErrTestSuite(t *testing.T) {
	suite.Run(t, new(ConflictErrTestSuite))
}

func (s *ConflictErrTestSuite) TestNewConflictError() {
	err := NewConflictError("some error")

	s.Error(err)
	s.IsType(ConflictError{}, err)

}

func (s *ConflictErrTestSuite) TestNewConflictErrorError() {
	err := NewConflictError("some error")
	s.Equal("some error", err.Error())
}
