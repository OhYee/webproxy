package proxy

import "testing"

func Test_getDomain(t *testing.T) {
	tests := []struct {
		url        string
		wantDomain string
		wantRest   string
	}{
		{"/github.com/OhYee/webproxy", "github.com", "/OhYee/webproxy"},
		{"/http:/github.com/OhYee/webproxy", "github.com", "/OhYee/webproxy"},
		{"/http://github.com/OhYee/webproxy", "github.com", "/OhYee/webproxy"},
		{"//http://github.com/OhYee/webproxy", "github.com", "/OhYee/webproxy"},
		{"/github.com", "github.com", "/"},
		{"/github.com/", "github.com", "/"},
		{"/github.com/OhYee", "github.com", "/OhYee"},
		{"/", "", "/"},
		{"/OhYee", "OhYee", "/"},
	}
	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			gotDomain, gotRest := getDomain(tt.url)
			if gotDomain != tt.wantDomain {
				t.Errorf("getDomain() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
			if gotRest != tt.wantRest {
				t.Errorf("getDomain() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
		})
	}
}
