package proxy

import (
	"fmt"
	"strings"
)

var (
	ignorePart = []string{"", "http:", "https:"}

	ignorePartMap = (func() map[string]struct{} {
		m := make(map[string]struct{})
		for _, part := range ignorePart {
			m[part] = struct{}{}
		}
		return m
	})()
)

func getDomain(url string) (domain string, rest string) {
	arr := strings.Split(url, "/")

	domain = ""
	for idx, item := range arr {
		if _, skip := ignorePartMap[item]; !skip {
			domain = item
			rest = strings.Join(arr[idx+1:], "/")
			break
		}
	}

	if len(rest) == 0 || rest[0] != '/' {
		rest = fmt.Sprintf("/%s", rest)
	}

	return
}
