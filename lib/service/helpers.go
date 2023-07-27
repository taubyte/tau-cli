package serviceLib

import (
	"github.com/taubyte/go-project-schema/project"
	"github.com/taubyte/go-project-schema/services"
	structureSpec "github.com/taubyte/go-specs/structure"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	"github.com/taubyte/utils/id"
)

type getter struct {
	project     project.Project
	application string
	service     services.Service
}

func get(name string) (info getter, err error) {
	info.project, info.application, err = applicationLib.SelectedProjectAndApp()
	if err != nil {
		return
	}

	info.service, err = info.project.Service(name, info.application)
	if err != nil {
		return
	}

	return
}

func list() (project project.Project, application string, services []string, err error) {
	project, application, err = applicationLib.SelectedProjectAndApp()
	if err != nil {
		return
	}

	local, global := project.Get().Services(application)
	if len(application) > 0 {
		services = local
	} else {
		services = global
	}

	return
}

func set(service *structureSpec.Service, new bool) error {
	info, err := get(service.Name)
	if err != nil {
		return err
	}

	if new == true {
		service.Id = id.Generate(info.project.Get().Id(), service.Name)
	}

	return info.service.SetWithStruct(true, service)
}
