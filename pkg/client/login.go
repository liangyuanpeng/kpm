package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"kcl-lang.io/kpm/pkg/reporter"
	"oras.land/oras-go/v2/registry/remote"
	remoteauth "oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/credentials"
)

// LoginOci will login to the oci registry.
func (c *KpmClient) LoginOci(hostname, username, password string) error {
	credCli, err := c.GetCredsClient()
	if err != nil {
		return err
	}

	registry, err := remote.NewRegistry(hostname)
	if err != nil {
		return err
	}
	cred := remoteauth.Credential{
		Username: username,
		Password: password,
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: c.insecureSkipTLSverify,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	registry.Client = &remoteauth.Client{Cache: remoteauth.NewCache(), Client: &http.Client{Transport: transport}}
	if hostname == "localhost:5001" {
		c.isPlainHttp = true
	}
	registry.PlainHTTP = c.isPlainHttp

	err = credentials.Login(context.Background(), credCli.Store, registry, cred)

	if err != nil {
		return reporter.NewErrorEvent(
			reporter.FailedLogin,
			err,
			fmt.Sprintf("failed to login '%s', please check registry, username and password is valid", hostname),
		)
	}

	return nil
}
