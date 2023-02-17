package repositoryLib

import authClient "github.com/taubyte/tau/singletons/auth_client"

func Register(repoID string) error {
	client, err := authClient.Load()
	if err != nil {
		return err
	}

	return client.RegisterRepository(repoID)
}
