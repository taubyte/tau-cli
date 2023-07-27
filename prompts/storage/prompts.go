package storagePrompts

import (
	storageFlags "github.com/taubyte/tau-cli/flags/storage"
	storageLib "github.com/taubyte/tau-cli/lib/storage"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func GetPublic(ctx *cli.Context, prev ...bool) bool {
	return prompts.GetOrAskForBoolDefaultTrue(ctx, storageFlags.Public.Name, PublicPrompt, prev...)
}

func GetVersioning(ctx *cli.Context, prev ...bool) bool {
	return prompts.GetOrAskForBool(ctx, storageFlags.Versioning.Name, VersioningPrompt, prev...)
}

func SelectABucket(ctx *cli.Context, prev ...string) string {
	return prompts.GetOrAskForSelection(ctx, storageFlags.BucketType.Name, BucketPrompt, storageLib.Buckets, prev...)
}

func GetOrRequireAMatch(ctx *cli.Context, prev ...string) string {
	return prompts.GetOrRequireAMatch(ctx, StorageMatch, prev...)
}
