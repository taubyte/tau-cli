package databasePrompts

import (
	"github.com/taubyte/go-project-schema/common"
	structureSpec "github.com/taubyte/go-specs/structure"
	databaseLib "github.com/taubyte/tau/lib/database"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) (*structureSpec.Database, error) {
	database := &structureSpec.Database{}

	taken, err := databaseLib.List()
	if err != nil {
		return nil, err
	}

	database.Name = prompts.GetOrRequireAUniqueName(ctx, NamePrompt, taken)
	database.Description = prompts.GetOrAskForADescription(ctx)
	database.Tags = prompts.GetOrAskForTags(ctx)

	database.Regex = prompts.GetMatchRegex(ctx)
	database.Match = GetOrRequireAMatch(ctx)
	database.Path = prompts.GetOrRequireAPath(ctx, prompts.PathPrompt)
	database.Local = prompts.GetOrAskForLocal(ctx)

	if GetEncryption(ctx) == true {
		database.Key = GetOrRequireAnEncryptionKey(ctx)
	}

	database.Min, database.Max, _, _ = GetOrAskForMinMax(ctx, 0, 0, true)

	database.Size, err = common.StringToUnits(prompts.GetSizeAndType(ctx, "", true))
	if err != nil {
		// TODO verbose
		return nil, err
	}

	return database, nil
}
