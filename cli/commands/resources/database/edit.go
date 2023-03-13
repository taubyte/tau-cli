package database

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	databaseFlags "github.com/taubyte/tau/flags/database"
	databaseI18n "github.com/taubyte/tau/i18n/database"
	databaseLib "github.com/taubyte/tau/lib/database"
	databasePrompts "github.com/taubyte/tau/prompts/database"
	databaseTable "github.com/taubyte/tau/table/database"
)

func (link) Edit() common.Command {
	return (&resources.Edit[*structureSpec.Database]{
		PromptsGetOrSelect: databasePrompts.GetOrSelect,
		PromptsEdit:        databasePrompts.Edit,
		TableConfirm:       databaseTable.Confirm,
		PromptsEditThis:    databasePrompts.EditThis,
		LibSet:             databaseLib.Set,
		I18nEdited:         databaseI18n.Edited,

		UniqueFlags: flags.Combine(
			flags.MatchRegex,
			flags.Match,
			flags.Local,
			databaseFlags.Encryption,
			databaseFlags.EncryptionKey,
			databaseFlags.Min,
			databaseFlags.Max,
			flags.Size,
			flags.SizeUnit,
		),
	}).Default()
}
