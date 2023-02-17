//go:build localAuthClient

package projectLib

import (
	"fmt"

	client "github.com/taubyte/go-auth-http"
)

var repoNum = 100000

func CreateRepository(client *client.Client, name, description string, private bool) (id string, err error) {
	id = fmt.Sprintf("%d", repoNum)
	repoNum += 1

	return
}
