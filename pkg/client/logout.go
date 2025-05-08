package client

import (
	"context"
	"fmt"

	"kcl-lang.io/kpm/pkg/reporter"
	"oras.land/oras-go/v2/registry/remote/credentials"
)

// LogoutOci will logout from the oci registry.
func (c *KpmClient) LogoutOci(hostname string) error {
	credCli, err := c.GetCredsClient()
	if err != nil {
		return err
	}

	err = credentials.Logout(context.Background(), credCli.Store, hostname)
	if err != nil {
		return reporter.NewErrorEvent(reporter.FailedLogout, err, fmt.Sprintf("failed to logout '%s'", hostname))
	}

	return nil
}
