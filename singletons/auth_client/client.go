package authClient

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/constants"
	"github.com/taubyte/tau-cli/env"
	"github.com/taubyte/tau-cli/i18n"
	networkI18n "github.com/taubyte/tau-cli/i18n/network"
	singletonsI18n "github.com/taubyte/tau-cli/i18n/singletons"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau-cli/singletons/dreamland"
	"github.com/taubyte/tau-cli/singletons/session"
	"github.com/taubyte/tau-cli/states"
	client "github.com/taubyte/tau/clients/http/auth"
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
		url = fmt.Sprintf("http://localhost:%d", getDreamlandAuthUrl())
	case common.RemoteNetwork:
		url = fmt.Sprintf("https://auth.tau.%s", profile.Network)
	case common.PythonTestNetwork:
		url = os.Getenv(constants.AuthURLEnvVarName)
		if url == "" {
			url = constants.ClientURL
		}
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

	client, err := client.New(
		states.Context,
		client.URL(url),
		client.Auth(profile.Token),
		client.Provider(profile.Provider),
	)
	if err != nil {
		return profile, nil, singletonsI18n.CreatingAuthClientFailed(err)
	}

	return profile, client, nil
}

func getDreamlandAuthUrl() int {
	ctx, ctxC := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxC()

	dreamClient, err := dreamland.Client(ctx)
	if err != nil {
		return 0
	}

	selectedUniverse, _ := env.GetCustomNetworkUrl()
	universe := dreamClient.Universe(selectedUniverse)
	echart, err := universe.Status()
	if err != nil {
		return 0
	}

	for _, node := range echart.Nodes {
		if strings.Contains(node.Name, "auth") {
			httpPort, ok := node.Value["http"]
			if !ok {
				return 0
			}

			return httpPort
		}
	}

	return 0
}
