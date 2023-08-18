package patrickClient

import (
	"context"
	"fmt"

	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/env"
	"github.com/taubyte/tau-cli/i18n"
	networkI18n "github.com/taubyte/tau-cli/i18n/network"
	singletonsI18n "github.com/taubyte/tau-cli/i18n/singletons"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau-cli/singletons/dreamland"
	"github.com/taubyte/tau-cli/singletons/session"
	"github.com/taubyte/tau-cli/states"
	"github.com/taubyte/tau/clients/http"
	client "github.com/taubyte/tau/clients/http/patrick"
)

var _client *client.Client

func Clear() {
	_client = nil
}

func getClientUrl() (url string, err error) {
	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return "", err
	}

	switch profile.NetworkType {
	case common.DreamlandNetwork:
		port, err := dreamland.HTTPPort(context.TODO(), "patrick")
		if err != nil {
			return "", err
		}

		url = fmt.Sprintf("http://127.0.0.1:%d", port)
	case common.RemoteNetwork:
		url = fmt.Sprintf("https://patrick.tau.%s", profile.Network)
	default:
		err = networkI18n.ErrorUnknownNetwork(profile.NetworkType)
	}

	return
}

func loadClient() (config.Profile, *client.Client, error) {
	profileName, exist := session.Get().ProfileName()
	if !exist {
		// Check for a default if no profiles are selected
		profileName, _, _ = loginLib.GetProfiles()
		if len(profileName) == 0 {
			i18n.Help().HaveYouLoggedIn()
			return config.Profile{}, nil, singletonsI18n.ProfileDoesNotExist()
		}
	}

	profile, err := config.Profiles().Get(profileName)
	if err != nil {
		return config.Profile{}, nil, err
	}

	selectedNetwork, _ := env.GetSelectedNetwork()
	if selectedNetwork == "" {
		i18n.Help().HaveYouSelectedANetwork()
		return config.Profile{}, nil, singletonsI18n.NoNetworkSelected()
	}

	url, err := getClientUrl()
	if err != nil {
		return config.Profile{}, nil, err
	}

	ops := []http.Option{http.URL(url), http.Auth(profile.Token)}
	client, err := client.New(states.Context, ops...)
	if err != nil {
		return profile, nil, singletonsI18n.CreatingPatrickClientFailed(err)
	}

	return profile, client, nil
}
