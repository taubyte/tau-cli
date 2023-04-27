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

func GetSelectedNetwork() (string, bool) {
	network, isSet := LookupEnv(constants.CurrentSelectedNetworkName)
	if isSet == true && len(network) > 0 {
		return network, isSet
	}

	return session.Get().SelectedNetwork()
}

func SetCustomNetworkUrl(c *cli.Context, network string) error {
	if justDisplayExport(c, constants.CustomNetworkUrlName, network) {
		return nil
	}

	return session.Set().CustomNetworkUrl(network)
}

func GetCustomNetworkUrl() (string, bool) {
	fqdn, isSet := LookupEnv(constants.CustomNetworkUrlName)
	if isSet == true && len(fqdn) > 0 {
		return fqdn, isSet
	}

	return session.Get().CustomNetworkUrl()
}
