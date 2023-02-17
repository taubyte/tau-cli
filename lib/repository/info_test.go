package repositoryLib_test

import (
	"fmt"
	"strconv"
	"testing"

	commonTest "github.com/taubyte/tau/common/test"
	repositoryLib "github.com/taubyte/tau/lib/repository"
	"github.com/taubyte/tau/singletons/session"
)

func TestInfo(t *testing.T) {
	err := session.Set().ProfileName("taubytetest")
	if err != nil {
		t.Error(err)
		return
	}

	info := &repositoryLib.Info{
		ID:   strconv.Itoa(commonTest.ConfigRepo.ID),
		Type: repositoryLib.WebsiteRepositoryType,
	}

	err = info.GetNameFromID()
	if err != nil {
		t.Error(err)
		return
	}

	expectedFullName := fmt.Sprintf("%s/%s", commonTest.GitUser, commonTest.ConfigRepo.Name)

	if info.FullName != expectedFullName {
		t.Errorf("Expected %s, got %s", expectedFullName, info.FullName)
		return
	}

	info.ID = ""
	err = info.GetIDFromName()
	if err != nil {
		t.Error(err)
		return
	}

	expectedID := commonTest.ConfigRepo.ID
	if info.ID != strconv.Itoa(expectedID) {
		t.Errorf("Expected %d, got %s", expectedID, info.ID)
		return
	}
}
