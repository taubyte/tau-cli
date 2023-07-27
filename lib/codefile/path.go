package codefile

import (
	"path"

	schemaCommon "github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/tau-cli/env"
	projectLib "github.com/taubyte/tau-cli/lib/project"
)

func Path(name, folder string) (CodePath, error) {
	projectConfig, err := projectLib.SelectedProjectConfig()
	if err != nil {
		return "", err
	}

	application, _ := env.GetSelectedApplication()

	var codePath string
	if len(application) > 0 {
		codePath = path.Join(projectConfig.CodeLoc(), schemaCommon.ApplicationFolder, application, folder, name)
	} else {
		codePath = path.Join(projectConfig.CodeLoc(), folder, name)
	}

	return CodePath(codePath), nil
}
