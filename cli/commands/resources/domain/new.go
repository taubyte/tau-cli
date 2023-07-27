package domain

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	domainFlags "github.com/taubyte/tau-cli/flags/domain"
	domainI18n "github.com/taubyte/tau-cli/i18n/domain"
	domainLib "github.com/taubyte/tau-cli/lib/domain"
	domainPrompts "github.com/taubyte/tau-cli/prompts/domain"
	domainTable "github.com/taubyte/tau-cli/table/domain"
)

func (link) New() common.Command {
	return (&resources.New[*structureSpec.Domain]{
		PromptsNew:        domainPrompts.New,
		TableConfirm:      domainTable.Confirm,
		PromptsCreateThis: domainPrompts.CreateThis,
		I18nCreated:       domainI18n.Created,
		UniqueFlags: flags.Combine(
			domainFlags.Generated,
			domainFlags.GeneratedPrefix,
			domainFlags.FQDN,
			domainFlags.CertType,
			domainFlags.Certificate,
			domainFlags.Key,
		),

		// Wrapping method to handle registration
		LibNew: func(resource *structureSpec.Domain) error {
			validator, err := domainLib.New(resource)
			if err != nil {
				return err
			}

			// Skipping registration check for generated FQDN
			isGeneratedFqdn, err := domainLib.IsAGeneratedFQDN(resource.Fqdn)
			if err != nil {
				return err
			}
			if isGeneratedFqdn {
				return nil
			}

			// Validate the fqdn provided
			clientResponse, err := validator.ValidateFQDN(resource.Fqdn)
			if err != nil {
				return err
			}

			domainTable.Registered(resource.Fqdn, clientResponse)
			return nil
		},
	}).Default()
}
