package env

import (
	"github.com/taubyte/tau/constants"
	"github.com/taubyte/tau/singletons/session"
	"github.com/urfave/cli/v2"
)

func SetSelectedNetwork(c *cli.Context, network string) error {
	if justDisplayExport(c, constants.CurrentSelectedNetworkName, network) {
		return nil
	}

	return session.Set().SelectedNetwork(network)
}

func GetSelectedNetwork() (string, error) {
	network, isSet := LookupEnv(constants.CurrentSelectedNetworkName)
	if isSet == true && len(network) > 0 {
		return network, nil
	}

	// Try to get profile from current session
	network, exist := session.Get().SelectedNetwork()
	if exist == true && len(network) > 0 {
		return network, nil
	}

	return network, nil
}

func SetCustomNetworkUrl(c *cli.Context, network string) error {
	if justDisplayExport(c, constants.CustomNetworkUrlName, network) {
		return nil
	}

	return session.Set().CustomNetworkUrl(network)
}

func GetCustomNetworkUrl() (string, error) {
	fqdn, isSet := LookupEnv(constants.CustomNetworkUrlName)
	if isSet == true && len(fqdn) > 0 {
		return fqdn, nil
	}

	// Try to get profile from current session
	fqdn, exist := session.Get().CustomNetworkUrl()
	if exist == true && len(fqdn) > 0 {
		return fqdn, nil
	}

	return fqdn, nil
}
