package manager

import "net/http"

type TransportItem struct {
	*http.Transport
	name string
}
