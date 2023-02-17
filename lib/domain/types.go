package domainLib

import client "github.com/taubyte/go-auth-http"

type Validator interface {
	ValidateFQDN(fqdn string) (response client.DomainResponse, err error)
}
