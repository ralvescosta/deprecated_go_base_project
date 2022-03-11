package valueObjects

type ValidateResult struct {
	IsValid bool
	Field   string
	Message string
}
