package proxy

import (
	"fmt"
	"net/http"

	"github.com/OhYee/webproxy/log"
)

func StartServer(addr string) error {
	return http.ListenAndServe(addr, &Handle{})
}

type Handle struct{}

func (h *Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domain, rest := getDomain(r.URL.Path)
	log.Infof("%s %s", domain, rest)

	f, ok := domainsMap[domain]
	if ok {
		f(domain, rest, w, r)
	} else {
		log.Errorf("%s not support, %#v", domain, r.URL)
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("%s not support\n", domain)))
	}
}
