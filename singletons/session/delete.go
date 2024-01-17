package session

import (
	"os"

	singletonsI18n "github.com/taubyte/tau-cli/i18n/singletons"
)

func Delete() error {
	processDir, found := nearestProcessDirectory(parentId())
	if !found || len(processDir) == 0 {
		return singletonsI18n.SessionNotFound()
	}

	err := os.RemoveAll(processDir)
	if err != nil {
		return singletonsI18n.SessionDeleteFailed(processDir, err)
	}

	return nil
}
