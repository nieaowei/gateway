package proxy_http

type ProxyError struct {
	msg  string
	Code int
}

func (l *ProxyError) Error() string {
	return l.msg
}

func (l *ProxyError) Is(x error) bool {
	return l.msg == x.Error()
}

func NewError(code int, msg string) *ProxyError {
	return &ProxyError{Code: code, msg: msg}
}

// middleware error start
const (
	Error_ServiceNotFound_Code = 3000 + iota
	Error_LBNotFound_Code
	Error_NoAvailableHost_Code
	Error_NoAvailableRedisService_Code
	Error_BlackListLimit_Code
	Error_WhiteListLimit_Code
	Error_NoAvailableTransport_Code
	Error_NoToken_Code
	Error_TokenInvalid_Code
	Error_NoAvailableApp_Code
)

var (
	Error_ServiceNotFound         = NewError(Error_ServiceNotFound_Code, "Service not found.")
	Error_LBNotFound              = NewError(Error_LBNotFound_Code, "Loadbalancer not found.")
	Error_NoAvailableHost         = NewError(Error_NoAvailableHost_Code, "No available host.")
	Error_NoAvailableRedisService = NewError(Error_NoAvailableRedisService_Code, "No available RedisService.")
	Error_BlackListLimit          = NewError(Error_BlackListLimit_Code, "Black list limit")
	Error_WhiteListLimit          = NewError(Error_WhiteListLimit_Code, "White list limit")
	Error_NoAvailableTransport    = NewError(Error_NoAvailableTransport_Code, "No available transport")
	Error_NoToken                 = NewError(Error_NoToken_Code, "No found Token")
	Error_TokenInvalid            = NewError(Error_TokenInvalid_Code, "Token is invalid")
	Error_NoAvailableApp          = NewError(Error_NoAvailableHost_Code, "no available app")
)

// middleware error end

// oauth error start
const (
	Error_AuthFormat_Code = 4000 + iota
	Error_AuthInfoFormat_Code
	Error_AppNotFound_Code
	Error_AppNotMatched_Code
)

var (
	Error_AuthFormat     = NewError(Error_AuthFormat_Code, "auth format error")
	Error_AuthInfoFormat = NewError(Error_AuthInfoFormat_Code, "auth info format error")
	Error_AppNotFound    = NewError(Error_AppNotFound_Code, "app info not found")
	Error_AppNotMatched  = NewError(Error_AppNotMatched_Code, "app info not matched")
)

// oauth error end
