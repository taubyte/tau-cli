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
	"github.com/urfave/cli/v2"
)

func (link) Edit() common.Command {
	var previousFQDN string

	return (&resources.Edit[*structureSpec.Domain]{
		I18nEdited:      domainI18n.Edited,
		PromptsEdit:     domainPrompts.Edit,
		TableConfirm:    domainTable.Confirm,
		PromptsEditThis: domainPrompts.EditThis,
		UniqueFlags: flags.Combine(
			domainFlags.FQDN,
			domainFlags.CertType,
			domainFlags.Certificate,
			domainFlags.Key,
		),

		// Wrapping methods to handle registration
		PromptsGetOrSelect: func(ctx *cli.Context) (*structureSpec.Domain, error) {
			resource, err := domainPrompts.GetOrSelect(ctx)
			if err != nil {
				return nil, err
			}

			previousFQDN = resource.Fqdn
			return resource, err
		},

		LibSet: func(resource *structureSpec.Domain) error {
			validator, err := domainLib.Set(resource)
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

			// Check if fqdn has changed and Validate the fqdn provided if it has
			if previousFQDN != resource.Fqdn {
				clientResponse, err := validator.ValidateFQDN(resource.Fqdn)
				if err != nil {
					return err
				}
				domainTable.Registered(resource.Fqdn, clientResponse)
			}

			return nil
		},
	}).Default()
}
