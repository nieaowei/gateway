package loadbalance

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
	Error_NoAvailableHost = NewError("No available host.")
	Error_AddNode         = NewError("AddHost node error.")
)
