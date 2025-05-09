package client

import (
	"context"
	"crypto/tls"
	"log"
	"net"
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
	// log.Println("c.insecureSkipTLSverify:", c.insecureSkipTLSverify)
	registry.Client = &auth.Client{Cache: auth.NewCache(), Client: &http.Client{Transport: transport}}
	// if c.insecureSkipTLSverify {
	// 	registry.PlainHTTP = true
	// }
	// if hostname == "localhost:5001" {
	// 	c.isPlainHttp = true
	// }
	// registry.PlainHTTP = c.isPlainHttp

	host, _, _ := net.SplitHostPort(hostname)
	// client.repo.PlainHTTP = false
	if host == "localhost" {
		// not specified, defaults to plain http for localhost
		registry.PlainHTTP = true
	}

	// If the plain http is specified in the settings file
	// Override the default value of the plain http
	if c.GetSettings() != nil {
		isPlainHttp, force := c.GetSettings().ForceOciPlainHttp()
		if force {
			registry.PlainHTTP = isPlainHttp
		}
	}

	err = credentials.Login(context.Background(), credCli.Store, registry, cred)

	if err != nil {
		return err
	}

	c1, err := credCli.Store.Get(context.TODO(), hostname)
	if err != nil {
		return err
	}
	log.Println("credCli.username:", c1.Username)

	return nil
}
