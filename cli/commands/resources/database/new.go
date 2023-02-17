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

func (link) New() common.Command {
	return (&resources.New[*structureSpec.Database]{
		PromptsNew:        databasePrompts.New,
		TableConfirm:      databaseTable.Confirm,
		PromptsCreateThis: databasePrompts.CreateThis,
		LibNew:            databaseLib.New,
		I18nCreated:       databaseI18n.Created,

		UniqueFlags: flags.Combine(
			flags.MatchRegex,
			flags.Match,
			flags.Path,
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
