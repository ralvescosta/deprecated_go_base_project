package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConflictError(t *testing.T) {
	t.Run("should create a conflictError correctly", func(t *testing.T) {
		err := NewConflictError("some error")

		assert.Error(t, err)
		assert.IsType(t, ConflictError{}, err)
	})
}

func Test_NewConflictError_Error(t *testing.T) {
	t.Run("should return error msg", func(t *testing.T) {
		err := NewConflictError("some error")

		assert.Equal(t, "some error", err.Error())
	})
}
