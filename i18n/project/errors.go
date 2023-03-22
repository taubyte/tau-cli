package projectI18n

import (
	"errors"
	"fmt"
)

const (
	bothFlagsCannotBeTrue         = "both %s and %s flags cannot be true"
	selectingVisibilityFailed     = "selecting visibility failed with: %s"
	gettingProjectsFailed         = "getting projects from auth failed with: %s"
	selectingAProjectPromptFailed = "selecting a project prompt failed with: %s"
	gettingRepositoriesFailed     = "getting repositories of `%s` failed with: %s"
	projectNotFound               = "project `%s` not found"
	gettingRepositoryURLsFailed   = "getting repository URLs for `%s` failed with: %s"
	cloningProjectFailed          = "cloning project `%s` failed with: %s"
	pullingProjectFailed          = "pulling project `%s` failed"
	pushingProjectFailed          = "pushing project `%s` failed"
	checkoutProjectFailed         = "checkout project `%s` failed"

	configRepoCreateFailed   = "creating config repository failed with: %s"
	codeRepoCreateFailed     = "creating code repository failed with: %s"
	configRepoRegisterFailed = "registering config repository failed with: %s"
	codeRepoRegisterFailed   = "registering code repository failed with: %s"
	creatingProjectFailed    = "creating project failed with: %s"
	projectBranchesNotEqual  = "config-`%s` and code-`%s` not on the same branch"

	ConfigRepo = "config repository: %s"
	CodeRepo   = "code repository: %s"
)

var (
	ProjectLocationEmpty     = errors.New("project location is empty")
	ConfigRepositoryNotFound = errors.New("config repository is not found")
	CodeRepositoryNotFound   = errors.New("code repository is not found")
)

func BothFlagsCannotBeTrue(flag1, flag2 string) error {
	return fmt.Errorf(bothFlagsCannotBeTrue, flag1, flag2)
}

func SelectingVisibilityFailed(err error) error {
	return fmt.Errorf(selectingVisibilityFailed, err)
}

func GettingProjectsFailed(err error) error {
	return fmt.Errorf(gettingProjectsFailed, err)
}

func SelectingAProjectPromptFailed(err error) error {
	return fmt.Errorf(selectingAProjectPromptFailed, err)
}

func GettingRepositoriesFailed(project string, err error) error {
	return fmt.Errorf(gettingRepositoriesFailed, project, err)
}

func ProjectNotFound(project string) error {
	return fmt.Errorf(projectNotFound, project)
}

func GettingRepositoryURLsFailed(project string, err error) error {
	return fmt.Errorf(gettingRepositoryURLsFailed, project, err)
}

func CloningProjectFailed(project string, err error) error {
	return fmt.Errorf(cloningProjectFailed, project, err)
}

func PullingProjectFailed(project string) error {
	return fmt.Errorf(pullingProjectFailed, project)
}

func CheckingOutProjectFailed(project string) error {
	return fmt.Errorf(checkoutProjectFailed, project)
}

func PushingProjectFailed(project string) error {
	return fmt.Errorf(pushingProjectFailed, project)
}

func ConfigRepoCreateFailed(err error) error {
	return fmt.Errorf(configRepoCreateFailed, err)
}

func CodeRepoCreateFailed(err error) error {
	return fmt.Errorf(codeRepoCreateFailed, err)
}

func ConfigRepoRegisterFailed(err error) error {
	return fmt.Errorf(configRepoRegisterFailed, err)
}

func CodeRepoRegisterFailed(err error) error {
	return fmt.Errorf(codeRepoRegisterFailed, err)
}

func CreatingProjectFailed(err error) error {
	return fmt.Errorf(creatingProjectFailed, err)
}

func ProjectBranchesNotEqual(branch1, branch2 string) error {
	return fmt.Errorf(projectBranchesNotEqual, branch1, branch2)
}
