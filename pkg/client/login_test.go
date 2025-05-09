package client

import "testing"

// go test -timeout 30s -run ^TestLoginOci$ kcl-lang.io/kpm/pkg/client -v
func TestLoginOci(t *testing.T) {
	// c := KpmClient{
	// 	insecureSkipTLSverify: true,
	// }
	// //192.168.3.187:5001
	// hostname := "ghcr.io"
	// name := "123"
	// password := ""
	// err := c.LoginOci(hostname, name, password)
	// if err != nil {
	// 	t.Error(err)
	// }
}

// go test -timeout 30s -run ^TestLogoutOci$ kcl-lang.io/kpm/pkg/client -v
func TestLogoutOci(t *testing.T) {
	c := KpmClient{
		// insecureSkipTLSverify: true,
	}
	//192.168.3.187:5001
	hostname := "invalid_registry"
	err := c.LogoutOci(hostname)
	if err != nil {
		t.Error(err)
	}
}
