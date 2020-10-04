package dao

import (
	"net/url"
	"strconv"
	"strings"
)

type LoadType int

const (
	Load_HTTP LoadType = iota
	Load_TCP
	Load_GRPC
)

type RoundType int

const (
	Round_Random RoundType = iota
	Round_RoudRobin
	Round_WeightRound
	Round_IpHash
)

type HttpRuleType int

const (
	HttpRule_PrefixURL HttpRuleType = iota
	HttpRule_Domain
)

type HeaderTransformType string

type HeaderTransformItem struct {
	Op  HeaderTransformOperationType
	Key string
	Val string
}

func (t *HeaderTransformType) GetSlice() []HeaderTransformItem {
	data := strings.Split(string(*t), "\n")
	out := []HeaderTransformItem{}
	for _, datum := range data {
		ss := strings.Split(datum, " ")
		item := HeaderTransformItem{
			Op:  HeaderTransformOperationType(ss[0]),
			Key: ss[1],
		}
		switch item.Op {
		case HeaderTransformOperation_Add, HeaderTransformOperation_Edit:
			item.Val = ss[2]
		}
		out = append(out, item)
	}
	return out
}

type HeaderTransformOperationType string

const (
	HeaderTransformOperation_Add  HeaderTransformOperationType = "add"
	HeaderTransformOperation_Edit HeaderTransformOperationType = "edit"
	HeaderTransformOperation_Del  HeaderTransformOperationType = "del"
)

type NeedHttpsType uint8

const (
	NeedHttps_Close NeedHttpsType = 0
	NeedHttps_Open  NeedHttpsType = 1
)

type NeedStripUriType uint8

const (
	NeedStripUri_Close NeedStripUriType = 0
	NeedStripUri_Open  NeedStripUriType = 1
)

type NeedWebsocketType uint8

const (
	NeedWebsocket_Close NeedWebsocketType = 0
	NeedWebsocket_Open  NeedWebsocketType = 1
)

type OpenAuthType uint8

const (
	OpenAuth_Close OpenAuthType = 0
	OpenAuth_Open  OpenAuthType = 1
)

type WeightListType string

type WeightListItem int

func (i *WeightListType) GetSlice() []WeightListItem {
	data := strings.Split(string(*i), "\n")
	out := []WeightListItem{}
	for _, datum := range data {
		num, _ := strconv.Atoi(datum)
		out = append(out, WeightListItem(num))
	}
	return out
}

type IpListType string

type IpListItem struct {
	url.URL
}

func (i *IpListType) GetSlice() []IpListItem {
	data := strings.Split(string(*i), "\n")
	out := []IpListItem{}
	for _, datum := range data {
		out = append(out, IpListItem{url.URL{Host: datum}})
	}
	return out
}

func (i *IpListItem) Url() url.URL {
	return i.URL
}
