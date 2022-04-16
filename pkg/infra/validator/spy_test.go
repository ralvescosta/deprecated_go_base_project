package validator

import (
	"testing"

	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateStructSpy(t *testing.T) {
	t.Run("execute correctly", func(t *testing.T) {
		v := NewValidatorSpy()

		type something struct {
			Id string `json:"id" validate:"required,uuid4"`
		}
		v.On("ValidateStruct", something{}).Return([]valueObjects.ValidateResult{})

		result := v.ValidateStruct(something{})

		assert.NotNil(t, result)
	})
}
