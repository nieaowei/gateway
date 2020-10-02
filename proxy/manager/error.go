package manager

type MgrError struct {
	msg string
}

func (l *MgrError) Error() string {
	return l.msg
}

func (l *MgrError) Is(x error) bool {
	return l.msg == x.Error()
}

func NewError(msg string) error {
	return &MgrError{msg: msg}
}

var (
	Error_NoMatchedService = NewError("No matched service.")
)
