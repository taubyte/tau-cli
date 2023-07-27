package domainLib

import (
	"github.com/taubyte/go-project-schema/domains"
	"github.com/taubyte/go-project-schema/project"
	structureSpec "github.com/taubyte/go-specs/structure"
	applicationLib "github.com/taubyte/tau-cli/lib/application"
	"github.com/taubyte/utils/id"
)

type getter struct {
	project     project.Project
	application string
	domain      domains.Domain
}

func get(name string) (info getter, err error) {
	info.project, info.application, err = applicationLib.SelectedProjectAndApp()
	if err != nil {
		return
	}

	info.domain, err = info.project.Domain(name, info.application)
	if err != nil {
		return
	}

	return
}

func list() (project project.Project, application string, domains []string, err error) {
	project, application, err = applicationLib.SelectedProjectAndApp()
	if err != nil {
		return
	}

	local, global := project.Get().Domains(application)
	if len(application) > 0 {
		domains = local
	} else {
		domains = global
	}

	return
}

func set(domain *structureSpec.Domain, new bool) (Validator, error) {
	info, err := get(domain.Name)
	if err != nil {
		return nil, err
	}

	if new == true {
		domain.Id = id.Generate(info.project.Get().Id(), domain.Name)
	}

	err = info.domain.SetWithStruct(true, domain)
	if err != nil {
		return nil, err
	}

	return newValidator(info), nil
}
