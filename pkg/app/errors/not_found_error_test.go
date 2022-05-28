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

func (s *NotFoundErrorTestSuite) TestCONFLICTNotFoundError() {
	err := NewNotFoundError("some error CONFLICT")

	s.Error(err)
	s.IsType(NotFoundError{}, err)
}

func (s *NotFoundErrorTestSuite) TestNotCONFLICTFoundErrorError() {
	//CONFLICT
	err := NewNotFoundError("some error")
	s.Equal("some error", err.Error())
}
