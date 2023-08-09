package dreamland

import (
	"context"
	"time"

	dreamland "github.com/taubyte/dreamland/service"
)

var dream_client *dreamland.Client

func Client(ctx context.Context) (*dreamland.Client, error) {
	if dream_client == nil {
		var err error
		dream_client, err = dreamland.New(ctx, dreamland.URL("http://localhost:1421"), dreamland.Timeout(15*time.Second))
		return dream_client, err
	}

	return dream_client, nil
}
