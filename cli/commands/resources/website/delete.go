package website

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	websiteI18n "github.com/taubyte/tau/i18n/website"
	websiteLib "github.com/taubyte/tau/lib/website"
	websitePrompts "github.com/taubyte/tau/prompts/website"
	websiteTable "github.com/taubyte/tau/table/website"
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
