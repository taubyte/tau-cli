package database

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	databaseFlags "github.com/taubyte/tau-cli/flags/database"
	databaseI18n "github.com/taubyte/tau-cli/i18n/database"
	databaseLib "github.com/taubyte/tau-cli/lib/database"
	databasePrompts "github.com/taubyte/tau-cli/prompts/database"
	databaseTable "github.com/taubyte/tau-cli/table/database"
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
