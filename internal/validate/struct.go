package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// StructError is the error returned by Struct.
type StructError struct {
	errs []error
}

// SetErrs sets the errors.
func (e *StructError) SetErrs(errs []error) {
	e.errs = errs
}

// AddErr adds an error.
func (e *StructError) AddErr(err error) {
	e.errs = append(e.errs, err)
}

// HasErrs returns true if there are errors.
func (e *StructError) HasErrs() bool {
	return len(e.errs) > 0
}

// Error returns all the errors as a string separated by commas.
func (e *StructError) Error() string {
	errStr := ""
	for idx, err := range e.errs {
		if idx > 0 {
			errStr += ", "
		}
		errStr += err.Error()
	}

	return errStr
}

// Errors returns all the errors as a slice of strings.
func (e *StructError) Errors() []string {
	errStrs := make([]string, len(e.errs))
	for idx, err := range e.errs {
		errStrs[idx] = err.Error()
	}

	return errStrs
}

// ErrorsRaw returns all the errors as a slice of errors.
func (e *StructError) ErrorsRaw() []error {
	return e.errs
}

// Struct validates the given struct using go-playground/validator.
func Struct[T any](sPointer *T) *StructError {
	err := validator.New().Struct(sPointer)

	if err != nil {
		errs := StructError{}

		if _, ok := err.(*validator.InvalidValidationError); ok {
			errs.AddErr(fmt.Errorf("validation error (check if it's a struct): %s", err))
		}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationError := range validationErrors {
				errs.AddErr(fmt.Errorf(
					"error in field %s: %s",
					validationError.StructField(),
					validationError.Tag(),
				))
			}
		}

		return &errs
	}

	return nil
}

// StructSlice validates the given slice of structs using go-playground/validator.
func StructSlice[T any](sPointerSlice *[]T) *StructError {
	for i, sPointer := range *sPointerSlice {
		num := i + 1
		if err := Struct(&sPointer); err != nil {
			se := &StructError{}
			errs := []error{fmt.Errorf("error in row %d", num)}
			errs = append(errs, err.ErrorsRaw()...)
			se.SetErrs(errs)
			return se
		}
	}

	return nil
}
