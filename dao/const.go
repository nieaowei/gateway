package dao

const (
	LoadType_HTTP = iota
	LoadType_TCP
	LoadType_GRPC
)

const (
	RoundType_Random = iota
	RoundType_RoudRobin
	RoundType_WeightRound
	RoundType_IpHash
)

const (
	DAO_False = iota
	DAO_True
)

const (
	HttpRuleType_PrefixURL = iota
	HttpRuleType_Domain
)
