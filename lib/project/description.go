package projectLib

import client "github.com/taubyte/go-auth-http"

func Description(p *client.Project) string {
	config, err := p.Config()
	if err != nil {
		return ""
	}

	return config.Description
}
