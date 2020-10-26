package dto

type DtoError struct {
	msg  string
	Code int
}

func (l *DtoError) Error() string {
	return l.msg
}

func (l *DtoError) Is(x error) bool {
	return l.msg == x.Error()
}

func NewError(code int, msg string) *DtoError {
	return &DtoError{Code: code, msg: msg}
}

const (
	//common
	ErrParamsValidCode = 5000 + iota

	//admin login
	ErrUserNotExistCode
	ErrPasswordCode
	// admin

)

var (
	//common
	ErrParamsValid = NewError(ErrParamsValidCode, "params format error")
	//admin login
	ErrUserNotExist = NewError(ErrUserNotExistCode, "user not exist")
	ErrPassword     = NewError(ErrPasswordCode, "password error")
	//admin

)
