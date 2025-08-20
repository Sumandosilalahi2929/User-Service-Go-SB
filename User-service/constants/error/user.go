package error

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrPasswordInCorrect     = errors.New("Password incorrect")
	ErrUsernameExist         = errors.New("username already exist")
	ErrPasswordDoestNotMatch = errors.New("password doest not match")
)

var userErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameExist,
	ErrPasswordDoestNotMatch,
}
