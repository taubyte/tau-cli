package authClient

import (
	"fmt"
	"os"

	client "github.com/taubyte/go-auth-http"
	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/constants"
	"github.com/taubyte/tau-cli/env"
	"github.com/taubyte/tau-cli/i18n"
	singletonsI18n "github.com/taubyte/tau-cli/i18n/singletons"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau-cli/singletons/session"
	"github.com/taubyte/tau-cli/states"
)

var _client *client.Client

func Clear() {
	_client = nil
}

func getClientUrl() (url string) {
	selectedNetwork, _ := env.GetSelectedNetwork()
	switch selectedNetwork {
	case common.DefaultNetwork:
		url = common.DefaultAuthUrl
	case common.DeprecatedNetwork:
		url = common.DeprecatedAuthUrl
	case common.CustomNetwork:
		customNetworkUrl, _ := env.GetCustomNetworkUrl()
		url = fmt.Sprintf("https://auth.tau.%s", customNetworkUrl)
	case common.PythonTestNetwork:
		url = os.Getenv(constants.AuthURLEnvVarName)
		if url == "" {
			url = constants.ClientURL
		}
	default:
		url = ""
	}

	return url
}

func loadClient() (config.Profile, *client.Client, error) {
	profileName, exist := session.Get().ProfileName()
	if exist == false {
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

	client, err := client.New(
		states.Context,
		client.URL(getClientUrl()),
		client.Auth(profile.Token),
		client.Provider(profile.Provider),
	)
	if err != nil {
		return profile, nil, singletonsI18n.CreatingAuthClientFailed(err)
	}

	return profile, client, nil
}
