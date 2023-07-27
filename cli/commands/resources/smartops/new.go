package smartops

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	smartopsI18n "github.com/taubyte/tau-cli/i18n/smartops"
	smartopsLib "github.com/taubyte/tau-cli/lib/smartops"
	smartopsPrompts "github.com/taubyte/tau-cli/prompts/smartops"
	smartopsTable "github.com/taubyte/tau-cli/table/smartops"
	"github.com/urfave/cli/v2"
)

func (link) New() common.Command {
	var templateURL string
	return (&resources.New[*structureSpec.SmartOp]{
		PromptsNew: func(ctx *cli.Context) (*structureSpec.SmartOp, error) {
			smartops, _templateURL, err := smartopsPrompts.New(ctx)
			templateURL = _templateURL
			return smartops, err
		},
		TableConfirm:      smartopsTable.Confirm,
		PromptsCreateThis: smartopsPrompts.CreateThis,
		LibNew: func(smartops *structureSpec.SmartOp) error {
			return smartopsLib.New(smartops, templateURL)
		},
		I18nCreated: smartopsI18n.Created,

		UniqueFlags: flags.Combine(
			flags.Timeout,
			flags.Memory,
			flags.MemoryUnit,
			flags.Source,
			flags.Call,
			flags.Template,
			flags.Language,
			flags.UseCodeTemplate,
		),
	}).Default()
}
