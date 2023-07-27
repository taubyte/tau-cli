package storage

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau-cli/cli/commands/resources/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/flags"
	storageFlags "github.com/taubyte/tau-cli/flags/storage"
	storageI18n "github.com/taubyte/tau-cli/i18n/storage"
	storageLib "github.com/taubyte/tau-cli/lib/storage"
	storagePrompts "github.com/taubyte/tau-cli/prompts/storage"
	storageTable "github.com/taubyte/tau-cli/table/storage"
)

func (link) Edit() common.Command {
	return (&resources.Edit[*structureSpec.Storage]{
		PromptsGetOrSelect: storagePrompts.GetOrSelect,
		PromptsEdit:        storagePrompts.Edit,
		TableConfirm:       storageTable.Confirm,
		PromptsEditThis:    storagePrompts.EditThis,
		LibSet:             storageLib.Set,
		I18nEdited:         storageI18n.Edited,

		UniqueFlags: flags.Combine(
			flags.MatchRegex,
			flags.Match,
			storageFlags.Public,
			flags.Size,
			flags.SizeUnit,
			storageFlags.BucketType,
			storageFlags.Versioning, // BucketType Object
			storageFlags.TTL,        // BucketType Streaming
		),
	}).Default()
}
