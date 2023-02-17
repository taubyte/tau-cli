package templates

import (
	git "github.com/taubyte/go-simple-git"
)

type templates struct {
	repository *git.Repository
}

type TemplateInfo struct {
	HideURL     bool
	URL         string
	Description string
}
