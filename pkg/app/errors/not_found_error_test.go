package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotFoundErrorTestSuite struct {
	suite.Suite
}

func TestNotFoundErrorTestSuite(t *testing.T) {
	suite.Run(t, new(NotFoundErrorTestSuite))
}

func (s *NotFoundErrorTestSuite) TestNotFoundErrorConflictHere() {
	err := NewNotFoundError("other conflict here")

	s.Error(err)
	s.IsType(NotFoundError{}, err)
}

func (s *NotFoundErrorTestSuite) TestNotFoundErrorError() {
	fmt.Println("WARNING")
	err := NewNotFoundError("some error")
	s.Equal("some error", err.Error())
}
