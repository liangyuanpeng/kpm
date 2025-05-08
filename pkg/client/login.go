package client

import (
	"context"

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
	err = credentials.Login(context.Background(), credCli.Store, registry, cred)

	if err != nil {
		return err
	}

	return nil
}
