package patrickClient

import (
	singletonsI18n "github.com/taubyte/tau-cli/i18n/singletons"
	patrickClient "github.com/taubyte/tau/clients/http/patrick"
)

func Load() (*patrickClient.Client, error) {
	if _client == nil {
		_, client, err := loadClient()
		if err != nil {
			return nil, singletonsI18n.LoadingAuthClientFailed(err)
		}

		_client = client
	}

	return _client, nil
}
