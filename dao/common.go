package dao

const (
	LoadTypeHttp = 0
	LoadTypeTcp  = 1
	LoadTypeGrpc = 2

	HttpRuleTypePrefixURL = 0
	HttpRuleTypeDomain    = 1
)

type PageSize struct {
	Size int
	No   int
	Info string
}
