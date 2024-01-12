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
