package pkg

import "errors"

var ErrInfraction = errors.New("infraction")

type Violation struct {
	Message string
}

func (v *Violation) Error() string {
	return v.Message
}

func (v *Violation) Is(target error) bool {
	return errors.Is(ErrInfraction, target)
}

var (
	ErrNotFound   = errors.New("entity not found")
	ErrEmptyField = errors.New("field is empty")
)

type EmptyFieldError struct {
	Field string
}

func (e *EmptyFieldError) Error() string {
	return e.Field + " cannot be empty"
}

func (e *EmptyFieldError) field() string { return e.Field }

func (e *EmptyFieldError) Is(err error) bool {
	return errors.Is(err, ErrEmptyField)
}

var ErrEntityNotFound = errors.New("entity not found")
