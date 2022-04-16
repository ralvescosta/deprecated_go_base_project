package interfaces

import (
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
)

type IValidator interface {
	ValidateStruct(data interface{}) []valueObjects.ValidateResult
}
