package constants

import (
	"fmt"
	"os"
	"path"
)

var (
	TauConfigFileName string
)

func init() {
	TauConfigFileName = os.Getenv(TauConfigFileNameEnvVarName)
	if TauConfigFileName == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Errorf("Who are you anyways! Trying to find your home directory failed with %s", err.Error()))
		}

		TauConfigFileName = path.Join(home, "tau")
	}
}
