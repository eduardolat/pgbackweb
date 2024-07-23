package validate

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Field string `validate:"required"`
}

func TestValidateStruct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		s := TestStruct{
			Field: "value",
		}
		err := Struct(&s)

		assert.Nil(t, err)
		assert.Equal(t, "value", s.Field)
	})

	t.Run("Fail", func(t *testing.T) {
		s := TestStruct{
			Field: "",
		}
		err := Struct(&s)

		assert.NotNil(t, err)
		assert.IsType(t, &StructError{}, err)
	})
}

func TestStructSlice(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		s := []TestStruct{
			{Field: "value1"},
			{Field: "value2"},
		}
		err := StructSlice(&s)

		assert.Nil(t, err)
	})

	t.Run("Fail on row 1", func(t *testing.T) {
		s := []TestStruct{
			{Field: ""},
			{Field: "value2"},
		}
		err := StructSlice(&s)

		assert.NotNil(t, err)
		assert.IsType(t, &StructError{}, err)
		assert.Contains(t, err.Error(), "error in row 1")
	})

	t.Run("Fail on row 2", func(t *testing.T) {
		s := []TestStruct{
			{Field: "value1"},
			{Field: ""},
		}
		err := StructSlice(&s)

		assert.NotNil(t, err)
		assert.IsType(t, &StructError{}, err)
		assert.Contains(t, err.Error(), "error in row 2")
	})
}

func TestStructError(t *testing.T) {
	t.Run("Error method", func(t *testing.T) {
		err := &StructError{
			errs: []error{errors.New("error1"), errors.New("error2")},
		}
		assert.Equal(t, "error1, error2", err.Error())
	})

	t.Run("Errors method", func(t *testing.T) {
		err := &StructError{
			errs: []error{errors.New("error1"), errors.New("error2")},
		}
		assert.Equal(t, []string{"error1", "error2"}, err.Errors())
	})

	t.Run("ErrorsRaw method", func(t *testing.T) {
		err := &StructError{
			errs: []error{errors.New("error1"), errors.New("error2")},
		}
		assert.Equal(
			t,
			[]error{errors.New("error1"), errors.New("error2")},
			err.ErrorsRaw(),
		)
	})

	t.Run("AddErr method", func(t *testing.T) {
		err := &StructError{}
		err.AddErr(errors.New("error1"))
		err.AddErr(errors.New("error2"))
		assert.Equal(t, []string{"error1", "error2"}, err.Errors())
	})

	t.Run("SetErrs method", func(t *testing.T) {
		err := &StructError{}
		err.SetErrs([]error{errors.New("error1"), errors.New("error2")})
		assert.Equal(t, []string{"error1", "error2"}, err.Errors())
	})

	t.Run("SetErrs method (overwrite)", func(t *testing.T) {
		err := &StructError{
			errs: []error{errors.New("error0")},
		}
		err.SetErrs([]error{errors.New("error1"), errors.New("error2")})
		assert.Equal(t, []string{"error1", "error2"}, err.Errors())
	})

	t.Run("HasErrs method (with errors)", func(t *testing.T) {
		err := &StructError{
			errs: []error{errors.New("error1"), errors.New("error2")},
		}
		assert.True(t, err.HasErrs())
	})

	t.Run("HasErrs method (without errors)", func(t *testing.T) {
		err := &StructError{}
		assert.False(t, err.HasErrs())
	})
}
