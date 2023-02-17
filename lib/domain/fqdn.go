package domainLib

import (
	"fmt"
	"strings"

	client "github.com/taubyte/go-auth-http"
	domainI18n "github.com/taubyte/tau/i18n/domain"
	projectLib "github.com/taubyte/tau/lib/project"
	authClient "github.com/taubyte/tau/singletons/auth_client"
)

type validator struct {
	getter
}

func NewValidator(name string) (Validator, error) {
	info, err := get(name)
	if err != nil {
		return nil, err
	}

	return &validator{info}, nil
}

// Internal does not require info
// TODO info should be global
func newValidator(info getter) Validator {
	return &validator{info}
}

func (r *validator) ValidateFQDN(fqdn string) (response client.DomainResponse, err error) {
	client, err := authClient.Load()
	if err != nil {
		return
	}

	return client.RegisterDomain(fqdn, r.project.Get().Id())
}

func NewGeneratedFQDN(prefix string) (string, error) {
	project, err := projectLib.SelectedProjectInterface()
	if err != nil {
		return "", err
	}

	// Get last eight characters of project id for use in fqdn
	projectID := project.Get().Id()
	if len(projectID) < 8 {
		return "", domainI18n.InvalidProjectIDEight(projectID)
	}
	projectID = strings.ToLower(projectID[len(projectID)-8:])

	// Generate fqdn
	fqdn := fmt.Sprintf("%s%d%s", projectID, ProjectDomainCount(project), GeneratedFqdnSuffix)

	// Attach prefix
	if len(prefix) > 0 {
		fqdn = fmt.Sprintf("%s-%s", prefix, fqdn)
	}

	return fqdn, nil
}

func IsAGeneratedFQDN(fqdn string) bool {
	return strings.HasSuffix(fqdn, GeneratedFqdnSuffix)
}
