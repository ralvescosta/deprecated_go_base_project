package validator

import (
	"testing"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateStruct(t *testing.T) {
	t.Run("should return validate erros when struct fields are wrong", func(t *testing.T) {
		sut := makeValidatorSutRtn()

		result := sut.validator.ValidateStruct(sut.wrong)

		expectedResult := []valueObjects.ValidateResult{
			{
				IsValid: false,
				Field:   "ID",
				Message: "ID invalid uuid4",
			},
			{
				IsValid: false,
				Field:   "Name",
				Message: "Name is required",
			},
		}
		assert.Len(t, result, 2)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return nil when don't have error", func(t *testing.T) {
		sut := makeValidatorSutRtn()

		result := sut.validator.ValidateStruct(sut.ok)

		assert.Nil(t, result)
	})
}

type someStruct struct {
	ID   string `json:"id" validate:"required,uuid4"`
	Name string `json:"name" validate:"required"`
}
type validatorSutRtn struct {
	validator interfaces.IValidator
	wrong     someStruct
	ok        someStruct
}

func makeValidatorSutRtn() validatorSutRtn {
	validator := NewValidator()

	wrong := someStruct{ID: "wrong_id"}
	ok := someStruct{ID: "a9c1a08b-4558-41d1-b69c-83fbb187e118", Name: "name"}

	return validatorSutRtn{validator, wrong, ok}
}
