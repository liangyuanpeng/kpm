package client

import "testing"

// go test -timeout 30s -run ^TestLoginOci$ kcl-lang.io/kpm/pkg/client -v
func TestLoginOci(t *testing.T) {
	c := KpmClient{
		insecureSkipTLSverify: true,
	}
	err := c.LoginOci("192.168.3.187:5001", "test", "1234")
	if err != nil {
		t.Error(err)
	}
}
