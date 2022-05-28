package errors

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotFoundErrorTestSuite struct {
	suite.Suite
}

func TestNotFoundErrorTestSuite(t *testing.T) {
	suite.Run(t, new(NotFoundErrorTestSuite))
}

func (s *NotFoundErrorTestSuite) TestNotFoundError() {
	err := NewNotFoundError("some error")

	s.Error(err)
	s.IsType(NotFoundError{}, err)
}

func (s *NotFoundErrorTestSuite) TestNotFoundErrorError() {
	err := NewNotFoundError("some error")
	s.Equal("some error", err.Error())
}
