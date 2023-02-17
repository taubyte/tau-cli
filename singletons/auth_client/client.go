package authClient

import (
	"os"

	client "github.com/taubyte/go-auth-http"
	"github.com/taubyte/tau/constants"
	"github.com/taubyte/tau/i18n"
	singletonsI18n "github.com/taubyte/tau/i18n/singletons"
	loginLib "github.com/taubyte/tau/lib/login"
	"github.com/taubyte/tau/singletons/config"
	"github.com/taubyte/tau/singletons/session"
	"github.com/taubyte/tau/states"
)

var _client *client.Client

func Clear() {
	_client = nil
}

func getClientUrl() (url string) {
	url = os.Getenv(constants.AuthURLEnvVarName)
	if url == "" {
		url = constants.ClientURL
	}

	return
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
