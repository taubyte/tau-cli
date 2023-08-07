package i18n

import (
	"errors"
	"fmt"
)

const (
	appCrashed      = "command failed with: %s"
	appCreateFailed = "creating new app failed with: %s"

	gettingCwdFailed = "getting current working directory failed with: %s"
)

func AppCrashed(err error) error {
	return fmt.Errorf(appCrashed, err)
}

func AppCreateFailed(err error) error {
	return fmt.Errorf(appCreateFailed, err)
}

func GettingCwdFailed(err error) error {
	return fmt.Errorf(gettingCwdFailed, err)
}

var ErrorTime0Invalid = errors.New("0 time is invalid")
