package loadbalance

type LoadBalanceError struct {
	msg string
}

func (l *LoadBalanceError) Error() string {
	return l.msg
}

func (l *LoadBalanceError) Is(x error) bool {
	return l.msg == x.Error()
}

func NewError(msg string) error {
	return &LoadBalanceError{msg: msg}
}

var (
	Error_NoAvailableHost = NewError("No available host.")
	Error_AddNode         = NewError("Add node error.")
)
