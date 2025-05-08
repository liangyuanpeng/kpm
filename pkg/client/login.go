package client

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
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
	log.Println("c.insecureSkipTLSverify:", c.insecureSkipTLSverify)
	registry.Client = &auth.Client{Cache: auth.NewCache(), Client: &http.Client{Transport: transport}}
	if !c.insecureSkipTLSverify {
		registry.PlainHTTP = true
	}

	err = credentials.Login(context.Background(), credCli.Store, registry, cred)

	if err != nil {
		return err
	}

	c1, err := credCli.Store.Get(context.TODO(), hostname)
	if err != nil {
		return err
	}
	log.Println("credCli:", c1.Username)

	return nil
}
