package viewmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrorViewModel_StringToErrorResponse(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := StringToErrorResponse("some error message")

		assert.IsType(t, ErrorMessage{}, sut)
	})
}
