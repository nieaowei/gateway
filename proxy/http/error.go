package proxy_http

type LoadBalanceError struct {
	msg  string
	Code int
}

func (l *LoadBalanceError) Error() string {
	return l.msg
}

func (l *LoadBalanceError) Is(x error) bool {
	return l.msg == x.Error()
}

func NewError(code int, msg string) *LoadBalanceError {
	return &LoadBalanceError{Code: code, msg: msg}
}

const (
	Error_ServiceNotFound_Code = 3000 + iota
	Error_LBNotFound_Code
	Error_NoAvailableHost_Code
	Error_NoAvailableRedisService_Code
	Error_BlackListLimit_Code
	Error_WhiteListLimit_Code
	Error_NoAvailableTransport_Code
)

var (
	Error_ServiceNotFound         = NewError(Error_ServiceNotFound_Code, "Service not found.")
	Error_LBNotFound              = NewError(Error_LBNotFound_Code, "Loadbalancer not found.")
	Error_NoAvailableHost         = NewError(Error_NoAvailableHost_Code, "No available host.")
	Error_NoAvailableRedisService = NewError(Error_NoAvailableRedisService_Code, "No available RedisService.")
	Error_BlackListLimit          = NewError(Error_BlackListLimit_Code, "Black list limit")
	Error_WhiteListLimit          = NewError(Error_WhiteListLimit_Code, "White list limit")
	Error_NoAvailableTransport    = NewError(Error_NoAvailableTransport_Code, "No available transport")
)
