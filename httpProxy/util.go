package httpProxy

import (
	"net/http"
	"strings"
)

func GetUrl(r *http.Request, host string) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	if host == "" {
		host = r.Host
	}
	url := strings.Join([]string{scheme, host, r.RequestURI}, "")
	return url
}
