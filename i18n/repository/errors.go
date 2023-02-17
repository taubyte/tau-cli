package repositoryI18n

import "fmt"

var (
	registeringRepositoryFailed = "Registering `%s` failed with: %s"
	unknownTemplate             = "unknown template `%s` must be one of `%v`"
)

func RegisteringRepositoryFailed(repo string, err error) error {
	return fmt.Errorf(registeringRepositoryFailed, repo, err)
}

func UnknownTemplate(selectedTemplate string, templates []string) error {
	return fmt.Errorf(unknownTemplate, selectedTemplate, templates)
}
