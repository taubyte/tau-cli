package session

import "github.com/taubyte/tau/constants"

type setter struct{}

func Set() Setter {
	getOrCreateSession()

	return setter{}
}

func (setter) ProfileName(value string) (err error) {
	return setKey(constants.CurrentSelectedProfileNameEnvVarName, value)
}

func (setter) SelectedProject(value string) (err error) {
	return setKey(constants.CurrentProjectEnvVarName, value)
}

func (setter) SelectedApplication(value string) (err error) {
	return setKey(constants.CurrentApplicationEnvVarName, value)
}
