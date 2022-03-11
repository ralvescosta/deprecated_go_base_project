package interfaces

import (
	valueObjects "markets/pkg/domain/value_objects"
)

type IValidator interface {
	ValidateStruct(data interface{}) []valueObjects.ValidateResult
}
