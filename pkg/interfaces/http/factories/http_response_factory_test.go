package factories

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	mErrors "markets/pkg/app/errors"
)

func Test_New(t *testing.T) {
	t.Run("should return new instance correctly", func(t *testing.T) {
		sut := NewHttpResponseFactory()

		assert.IsType(t, HttpResponseFactory{}, sut)
	})
}

func Test_Ok(t *testing.T) {
	t.Run("should return httpStatus 200", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.Ok(nil, nil).StatusCode, http.StatusOK)
	})
}

func Test_Created(t *testing.T) {
	t.Run("should return httpStatus 201", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.Created(nil, nil).StatusCode, http.StatusCreated)
	})
}

func Test_NoConted(t *testing.T) {
	t.Run("should return httpStatus 204", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.NoContent(nil).StatusCode, http.StatusNoContent)
	})
}

func Test_BadRequest(t *testing.T) {
	t.Run("should return httpStatus 400", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.BadRequest("", nil).StatusCode, http.StatusBadRequest)
	})
}

func Test_Unauthorized(t *testing.T) {
	t.Run("should return httpStatus 401", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.Unauthorized("", nil).StatusCode, http.StatusUnauthorized)
	})
}

func Test_Forbidden(t *testing.T) {
	t.Run("should return httpStatus 403", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.Forbidden("", nil).StatusCode, http.StatusForbidden)
	})
}

func Test_NotFound(t *testing.T) {
	t.Run("should return httpStatus 404", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.NotFound("", nil).StatusCode, http.StatusNotFound)
	})
}

func Test_Conflict(t *testing.T) {
	t.Run("should return httpStatus 409", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.Conflict("", nil).StatusCode, http.StatusConflict)
	})
}

func Test_InternalServerError(t *testing.T) {
	t.Run("should return httpStatus 500", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.InternalServerError("", nil).StatusCode, http.StatusInternalServerError)
	})
}

func Test_GenericResponse(t *testing.T) {
	t.Run("should return httpStatus 200", func(t *testing.T) {
		sut := HttpResponseFactory{}

		assert.Equal(t, sut.GenericResponse(http.StatusOK, "", nil).StatusCode, http.StatusOK)
	})
}

func Test_ErrorResponseMapper(t *testing.T) {
	t.Run("should map notFoundError to NotFound response", func(t *testing.T) {
		err := mErrors.NewNotFoundError("some error")
		sut := HttpResponseFactory{}

		result := sut.ErrorResponseMapper(err, nil)

		assert.Equal(t, result.StatusCode, http.StatusNotFound)
	})

	t.Run("should map conflictError to Conflict response", func(t *testing.T) {
		err := mErrors.NewConflictError("some error")
		sut := HttpResponseFactory{}

		result := sut.ErrorResponseMapper(err, nil)

		assert.Equal(t, result.StatusCode, http.StatusConflict)
	})

	t.Run("should map unmapped error to InternalServerError response", func(t *testing.T) {
		err := errors.New("some error")
		sut := HttpResponseFactory{}

		result := sut.ErrorResponseMapper(err, nil)

		assert.Equal(t, result.StatusCode, http.StatusInternalServerError)
	})
}
