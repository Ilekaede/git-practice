package errors

import (
	"errors"
	"fmt"
)

var(
	ErrAlreadyRegistered = newErrDomain("ErrAlreadyRegistered", "you have been already registered")
	ErrInvalidEmail = newErrDomain("ErrInvalidEmail", "invalid email address")
	ErrPasswordMismatch = newErrDomain("ErrPasswordMismatch", "the password is incorrect")
	ErrPasswordTooShort = newErrDomain("ErrPasswordTooShort", "password is too short")
	ErrNotFoundUser = newErrDomain("ErrNotFoundUser", "user not found")
)

var (
	ErrNotFoundTask = newErrDomain("ErrNotFoundTask", "task not found")
	ErrContentEmpty           = newErrDomain("ErrContentEmpty", "Do not empty the content")
	ErrInvalidTaskState       = newErrDomain("ErrInvalidTaskState", "invalid task state , please select todo/doing/done")
	ErrForbiddenTaskOperation = newErrDomain("ErrForbiddenTaskOperation", "can't operate others tasks")
)

type ErrDomain struct{
	err error
}

func IsDomainErr(target error) bool{
	var errDomain *ErrDomain
	return errors.As(target, &errDomain)
}

func (e *ErrDomain) Error() string{
	return e.err.Error()
}


func Is(err, target error) bool{
	return errors.Is(err, target)
}
func New(message string) error{
	return errors.New(message)
}
