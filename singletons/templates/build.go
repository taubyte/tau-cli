package templates

import (
	gosimplegit "github.com/taubyte/go-simple-git"
)

func Get() *templates {
	getOrCreateTemplates()

	return _templates
}

func Repository() *gosimplegit.Repository {
	return Get().repository
}

type templateYaml struct {
	// parameters must be exported for the yaml parser
	Name        string
	Description string
	Icon        string
	URL         string
}
