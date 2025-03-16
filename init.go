package validation

func NewValidator(value any) *ValidationValue {
	return &ValidationValue{value: value}
}
