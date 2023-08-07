package network

import "errors"

const (
	flagError = "only set one flag corresponding to a network"
)

func FlagError() error {
	return errors.New(flagError)
}
