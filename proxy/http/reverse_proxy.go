package proxy_http

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

type UrlPathFormatHandler func(url2 *url.URL) string

func NewHttpProxy(target *url.URL, trans *http.Transport, format UrlPathFormatHandler) *httputil.ReverseProxy {
	//httputil.NewSingleHostReverseProxy()
	targetQuery := target.RawQuery

	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		if format == nil {
			format = func(url2 *url.URL) string {
				return url2.Path
			}
		}
		req.URL.Path = singleJoiningSlash(target.Path, format(req.URL))
		req.Host = target.Host

		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "nekilc")
		}
	}

	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	return &httputil.ReverseProxy{
		Transport:    trans,
		ErrorHandler: errFunc,
		Director:     director,
	}
}
