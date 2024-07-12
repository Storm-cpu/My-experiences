package server

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// custom variable
var (
	DocumentExts = []string{".pdf", ".jpg", ".jpeg", ".png"}
	ImageExts    = []string{".jpg", ".jpeg", ".png"}
)

// CustomValidator holds custom validator
type CustomValidator struct {
	V *validator.Validate
}

// NewValidator creates new custom validator
func NewValidator() *CustomValidator {
	V := validator.New()
	V.RegisterValidation("date", validateDate)
	V.RegisterValidation("mobile", validateMobile)
	return &CustomValidator{V}
}

// Validate validates the request
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.V.Struct(i)
}

func validateDate(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	re := regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}(T00:00:00Z)?$`)
	return re.MatchString(val)
}

func validateMobile(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	re := regexp.MustCompile(`^([0]?(3|5|7|8|9|1[2|6|8|9]))([0-9]{8})\b`)
	return re.MatchString(strings.Replace(val, " ", "", -1))
}