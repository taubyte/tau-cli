package storage

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	resources "github.com/taubyte/tau/cli/commands/resources/common"
	"github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/flags"
	storageFlags "github.com/taubyte/tau/flags/storage"
	storageI18n "github.com/taubyte/tau/i18n/storage"
	storageLib "github.com/taubyte/tau/lib/storage"
	storagePrompts "github.com/taubyte/tau/prompts/storage"
	storageTable "github.com/taubyte/tau/table/storage"
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
