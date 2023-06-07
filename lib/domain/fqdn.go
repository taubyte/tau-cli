package domainLib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	client "github.com/taubyte/go-auth-http"
	"github.com/taubyte/tau/common"
	"github.com/taubyte/tau/env"
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

	parseFqdn := func(suffix string) string {
		return fmt.Sprintf("%s%d%s", projectID, ProjectDomainCount(project), suffix)
	}

	// Generate fqdn
	var fqdn string
	selectedNetwork, _ := env.GetSelectedNetwork()

	switch selectedNetwork {
	case common.DefaultNetwork:
		fqdn = parseFqdn(GeneratedFqdnSuffix)
	case common.DeprecatedNetwork:
		fqdn = parseFqdn(DeprecatedGeneratedFqdnSuffix)
	case common.PythonTestNetwork:
		fqdn = parseFqdn(DeprecatedGeneratedFqdnSuffix)
	case common.CustomNetwork:
		customNetworkUrl, _ := env.GetCustomNetworkUrl()
		customGeneratedFqdn, err := FetchCustomNetworkGeneratedFqdn(customNetworkUrl)
		if err != nil {
			return "", err
		}

		fqdn = parseFqdn(customGeneratedFqdn)
	}

	// Attach prefix
	if len(prefix) > 0 {
		fqdn = fmt.Sprintf("%s-%s", prefix, fqdn)
	}

	return fqdn, nil
}

func IsAGeneratedFQDN(fqdn string) (bool, error) {
	selectedNetwork, _ := env.GetSelectedNetwork()
	switch selectedNetwork {
	case common.DeprecatedNetwork:
		return strings.HasSuffix(fqdn, DeprecatedGeneratedFqdnSuffix), nil
	case common.CustomNetwork:
		customNetworkUrl, _ := env.GetCustomNetworkUrl()
		customGeneratedFqdn, err := FetchCustomNetworkGeneratedFqdn(customNetworkUrl)
		if err != nil {
			return false, err
		}

		return strings.HasSuffix(fqdn, customGeneratedFqdn), nil
	default:
		return strings.HasSuffix(fqdn, GeneratedFqdnSuffix), nil
	}
}

// TODO: Move to specs
func FetchCustomNetworkGeneratedFqdn(fqdn string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://seer.tau.%s/network/config", fqdn))
	if err != nil {
		return "", fmt.Errorf("fetching generated url prefix for fqdn `%s` failed with: %s", fqdn, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response failed with: %s", err)
	}

	bodyStr := strings.Trim(string(body), "\"")

	return formatGeneratedSuffix(bodyStr), nil
}

func formatGeneratedSuffix(suffix string) string {
	if !strings.HasPrefix(suffix, ".") {
		suffix = "." + suffix
	}

	return suffix
}
