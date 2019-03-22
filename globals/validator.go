package globals

import "github.com/go-playground/validator"

// CustomValidator 自定义validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate 校验过程
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewDefaultValidator 返回默认校验器
func NewDefaultValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
