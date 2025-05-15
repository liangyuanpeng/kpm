package client

import "testing"

func TestLogout(t *testing.T) {
	kpmcli, err := NewKpmClient()
	if err != nil {
		// reporter.Fatal(err)
		panic(err)
	}
	kpmcli.LoginOci("172.88.0.8:5002", "test", "1234")
	kpmcli.LogoutOci("172.88.0.8:5002")
}
