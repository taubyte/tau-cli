package website

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	websiteI18n "github.com/taubyte/tau-cli/i18n/website"
	websiteLib "github.com/taubyte/tau-cli/lib/website"
	websitePrompts "github.com/taubyte/tau-cli/prompts/website"
	websiteTable "github.com/taubyte/tau-cli/table/website"
)

func (link) Delete() common.Command {
	return (&resources.Delete[*structureSpec.Website]{
		PromptsGetOrSelect: websitePrompts.GetOrSelect,
		TableConfirm:       websiteTable.Confirm,
		PromptsDeleteThis:  websitePrompts.DeleteThis,
		LibDelete:          websiteLib.Delete,
		I18nDeleted:        websiteI18n.Deleted,
	}).Default()
}
