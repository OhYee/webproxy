package proxy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/OhYee/webproxy/log"
)

var (
	proxyEnvs  = []string{"HTTP_PROXY", "http_proxy", "HTTPS_PROXY", "https_proxy"}
	domainsMap = map[string]func(domain, rest string, w http.ResponseWriter, r *http.Request){
		"github.com":                    reverseProxy,
		"api.github.com":                reverseProxy,
		"github.githubassets.com":       reverseProxy,
		"gist.githubusercontent.com":    reverseProxy,
		"avatars.githubusercontent.com": reverseProxy,
		"camo.githubusercontent.com":    reverseProxy,
		"gist.github.com":               gistProxy,
	}

	httpClient = http.DefaultClient

	serverHost = strings.TrimRight(os.Getenv("SERVER_HOST"), "/")

	schemeProcotol = func() string {
		scheme := os.Getenv("SCHEME")
		if scheme == "" {
			scheme = "https"
		}
		return scheme
	}()
)

func init() {
	var u *url.URL
	var err error
	for _, item := range proxyEnvs {
		if v := os.Getenv(item); v != "" {
			u, err = url.Parse(v)
			if err == nil && u != nil {
				break
			}
		}
	}

	if u != nil {
		httpClient.Transport = &http.Transport{
			Proxy: func(*http.Request) (*url.URL, error) {
				return u, nil
			},
		}
	}

}

func reverseProxy(domain, rest string, w http.ResponseWriter, r *http.Request) {
	log.Infof("reverseProxy %s => %s%s", r.URL.String(), domain, rest)

	u, err := url.Parse(fmt.Sprintf("%s://%s", schemeProcotol, domain))
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	rp := httputil.NewSingleHostReverseProxy(u)
	rp.Transport = httpClient.Transport
	rp.Director = func(r *http.Request) {
		r.RequestURI = ""
		r.Host = domain
		if r.URL.Scheme == "" {
			r.URL.Scheme = schemeProcotol
		}
		r.URL.Host = domain
		r.URL.Path = rest
	}
	rp.ServeHTTP(w, r)
}

func gistProxy(domain, rest string, w http.ResponseWriter, r *http.Request) {
	log.Infof("gistProxy %s%s", domain, rest)

	resp, err := httpClient.Get(fmt.Sprintf("https://gist.github.com%s", rest))
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	b = bytes.Replace(
		b,
		[]byte("https://github.githubassets.com"),
		[]byte(fmt.Sprintf("%s/github.githubassets.com", serverHost)),
		1,
	)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
