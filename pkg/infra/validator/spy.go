package validator

import (
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"

	"github.com/stretchr/testify/mock"
)

type ValidatorSpy struct {
	mock.Mock
}

func (pst ValidatorSpy) ValidateStruct(m interface{}) []valueObjects.ValidateResult {
	args := pst.Called(m)

	return args.Get(0).([]valueObjects.ValidateResult)
}

func NewValidatorSpy() *ValidatorSpy {
	return new(ValidatorSpy)
}
