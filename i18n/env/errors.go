package envI18n

import "errors"

var (
	UserNotFound    = errors.New("user not found")
	ProjectNotFound = errors.New("project not found")
)
