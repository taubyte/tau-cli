package repositoryLib

import (
	"fmt"

	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
)

func (info *Info) GetNameFromID() error {
	client, err := authClient.Load()
	if err != nil {
		return fmt.Errorf("loading auth client failed with: %w", err)
	}

	repo, err := client.GetRepositoryById(info.ID)
	if err != nil {
		return fmt.Errorf("getting repo by id failed with: %w", err)
	}

	info.FullName = repo.Get().FullName()
	if len(info.FullName) == 0 {
		return fmt.Errorf("could not find repository with id `%s`", info.ID)
	}

	return nil
}

func (info *Info) GetIDFromName() error {
	client, err := authClient.Load()
	if err != nil {
		return err
	}

	repo, err := client.GetRepositoryByName(info.FullName)
	if err != nil {
		return err
	}

	id := repo.Get().ID()
	if len(id) == 0 {
		return fmt.Errorf("could not find repository with name `%s`", info.FullName)
	}

	info.ID = id
	return nil
}
