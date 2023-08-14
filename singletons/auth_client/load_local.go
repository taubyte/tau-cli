//go:build localAuthClient

package authClient

import (
	singletonsI18n "github.com/taubyte/tau-cli/i18n/singletons"
	client "github.com/taubyte/tau/clients/http/auth"
)

func Load() (*client.Client, error) {
	if _client == nil {
		_, client, err := loadClient()
		if err != nil {
			return nil, singletonsI18n.LoadingAuthClientFailed(err)
		}

		_client = client
	}

	return _client, nil
}
