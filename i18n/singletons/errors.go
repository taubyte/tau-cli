package singletonsI18n

import (
	"fmt"
)

const (
	// Common
	creatingSeerAtLocFailed = "creating seer at `%s` failed with %s:"

	// Session
	sessionFileLocationEmpty  = "session file location is empty and could not discover or create"
	sessionSettingKeyFailed   = "setting session key `%s` to `%s` failed with: %s"
	sessionDeletingKeyFailed  = "deleting session key `%s`  failed with: %s"
	sessionDeleteFailed       = "deleting session at %s failed with: %s"
	sessionListFailed         = "getting session items failed with: %s"
	sessionNotFound           = "no session found"
	sessionCreateFailed       = "creating session file at `%s` failed with: %s"
	creatingSessionFileFailed = "creating session file failed with %s:"

	// Config
	creatingConfigFileFailed = "creating config file failed with %s:"
	gettingProfileFailedWith = "getting profile `%s` from config failed with %s:"
	settingProfileFailedWith = "setting profile `%s` in config failed with %s:"

	gettingProjectFailedWith   = "getting project `%s` from config failed with %s:"
	settingProjectFailedWith   = "setting project `%s` in config failed with %s:"
	deletingProjectFailedWith  = "deleting project `%s` from config failed with %s:"
	projectLocationNotFound    = "project `%s` location not found"
	openingProjectConfigFailed = "opening project config at `%s` failed with: %s"
	projectAlreadyCloned       = "project `%s` already cloned in: `%s`"

	// Auth_client
	profileDoesNotExist      = "profile does not exist"
	creatingAuthClientFailed = "creating auth client failed with: %s"
	loadingAuthClientFailed  = "loading auth client failed with: %s"

	//Network
	noNetworkSelected = "no network selected"
)

func CreatingSeerAtLocFailed(loc string, err error) error {
	return fmt.Errorf(creatingSeerAtLocFailed, loc, err)
}

func SessionFileLocationEmpty() error {
	return fmt.Errorf(sessionFileLocationEmpty)
}

func SessionSettingKeyFailed(key string, value interface{}, err error) error {
	return fmt.Errorf(sessionSettingKeyFailed, key, value, err)
}

func SessionDeletingKeyFailed(key string, err error) error {
	return fmt.Errorf(sessionDeletingKeyFailed, key, err)
}

func SessionDeleteFailed(loc string, err error) error {
	return fmt.Errorf(sessionDeleteFailed, loc, err)
}

func SessionListFailed(err error) error {
	return fmt.Errorf(sessionListFailed, err)
}

func SessionNotFound() error {
	return fmt.Errorf(sessionNotFound)
}

func SessionCreateFailed(loc string, err error) error {
	return fmt.Errorf(sessionCreateFailed, loc, err)
}

func CreatingSessionFileFailed(err error) error {
	return fmt.Errorf(creatingSessionFileFailed, err)
}

func CreatingConfigFileFailed(err error) error {
	return fmt.Errorf(creatingConfigFileFailed, err)
}

func GettingProfileFailedWith(profile string, err error) error {
	return fmt.Errorf(gettingProfileFailedWith, profile, err)
}

func SettingProfileFailedWith(profile string, err error) error {
	return fmt.Errorf(settingProfileFailedWith, profile, err)
}

func GettingProjectFailedWith(project string, err error) error {
	return fmt.Errorf(gettingProjectFailedWith, project, err)
}

func SettingProjectFailedWith(project string, err error) error {
	return fmt.Errorf(settingProjectFailedWith, project, err)
}

func DeletingProjectFailedWith(project string, err error) error {
	return fmt.Errorf(deletingProjectFailedWith, project, err)
}

func ProjectLocationNotFound(project string) error {
	return fmt.Errorf(projectLocationNotFound, project)
}

func OpeningProjectConfigFailed(loc string, err error) error {
	return fmt.Errorf(openingProjectConfigFailed, loc, err)
}

func ProjectAlreadyCloned(project, loc string) error {
	return fmt.Errorf(projectAlreadyCloned, project, loc)
}

func ProfileDoesNotExist() error {
	return fmt.Errorf(profileDoesNotExist)
}

func CreatingAuthClientFailed(err error) error {
	return fmt.Errorf(creatingAuthClientFailed, err)
}

func LoadingAuthClientFailed(err error) error {
	return fmt.Errorf(loadingAuthClientFailed, err)
}

func NoNetworkSelected() error {
	return fmt.Errorf(noNetworkSelected)
}
