package downloader

import (
	"context"
	"fmt"

	remoteauth "oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/credentials"
	orascredentials "oras.land/oras-go/v2/registry/remote/credentials"
)

// CredClient is the client to get the credentials.
type CredClient struct {
	Store *orascredentials.DynamicStore
}

// LoadCredentialFile loads the credential file and return the CredClient.
func LoadCredentialFile(filepath string) (*CredClient, error) {
	// authClient, err := dockerauth.NewClientWithDockerFallback(filepath)
	// if err != nil {
	// 	return nil, err
	// }
	// dockerAuthClient, ok := authClient.(*dockerauth.Client)
	// if !ok {
	// 	return nil, fmt.Errorf("authClient is not *docker.Client type")
	// }

	store, err := credentials.NewStore(filepath, credentials.StoreOptions{
		AllowPlaintextPut: true,
	})
	if err != nil {
		return nil, fmt.Errorf("store init failed!" + err.Error())
	}

	return &CredClient{
		Store: store,
	}, nil
}

// Credential will reture the credential info cache in CredClient
func (cred *CredClient) Credential(hostName string) (*remoteauth.Credential, error) {
	if len(hostName) == 0 {
		return nil, fmt.Errorf("hostName is empty")
	}
	remoteauth1, err := cred.Store.Get(context.Background(), hostName)
	// username, password, err := cred.credsClient.Credential(hostName)
	if err != nil {
		return nil, err
	}

	return &remoteauth.Credential{
		Username: remoteauth1.Username,
		Password: remoteauth1.Password,
	}, nil
}
