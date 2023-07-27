package authClient_test

import (
	"testing"

	commonTest "github.com/taubyte/tau-cli/common/test"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau-cli/singletons/session"
)

func TestClient(t *testing.T) {
	profiles := config.Profiles()
	testProfileName := "prof1"
	testProfile := config.Profile{
		Provider: "github",
		Token:    commonTest.GitToken(),
		Default:  false,
	}

	err := profiles.Set(testProfileName, testProfile)
	if err != nil {
		t.Error(err)
		return
	}

	err = session.Set().ProfileName(testProfileName)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = authClient.Load()
	if err != nil {
		t.Error(err)
		return
	}
}
