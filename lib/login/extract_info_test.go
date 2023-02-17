package loginLib

import (
	"testing"

	commonTest "github.com/taubyte/tau/common/test"
)

func TestInfo(t *testing.T) {
	name, email, err := extractInfo(commonTest.GitToken(), "github")
	if err != nil {
		t.Error(err)
		return
	}

	expectedName := "taubyte-test"
	expectedEmail := "taubytetest@gmail.com"

	if name != expectedName {
		t.Errorf("Expected name: %s, got: %s", expectedName, name)
	}
	if email != expectedEmail {
		t.Errorf("Expected email: %s, got: %s", expectedEmail, email)
	}
}
