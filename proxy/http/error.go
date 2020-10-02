package proxy_http

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
	Error_ServiceNotFound = NewError("Service not found.")
	Error_LBNotFound      = NewError("Loadbalancer not found.")
	Error_NoAvailableHost = NewError("No available host.")
)
