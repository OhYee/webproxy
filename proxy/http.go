package proxy

import (
	"net/http"
	"net/url"

	"github.com/OhYee/webproxy/log"
)

func StartHTTPServer(addr, redirectAddr string) error {
	return http.ListenAndServe(addr, &httpHandle{redirectAddr: redirectAddr})
}

type httpHandle struct {
	redirectAddr string
}

func (h *httpHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(h.redirectAddr)
	if err != nil {
		log.Errorf("can not parse domain, due to", err)
	}
	domain := u.Host
	rest := r.URL.Path
	log.Infof("%s %s", domain, rest)

	reverseProxy(domain, rest, w, r)
}
