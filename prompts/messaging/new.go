package messagingPrompts

import (
	structureSpec "github.com/taubyte/go-specs/structure"
	messagingLib "github.com/taubyte/tau-cli/lib/messaging"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func New(ctx *cli.Context) (*structureSpec.Messaging, error) {
	messaging := &structureSpec.Messaging{}

	taken, err := messagingLib.List()
	if err != nil {
		return nil, err
	}

	messaging.Name = prompts.GetOrRequireAUniqueName(ctx, NamePrompt, taken)
	messaging.Description = prompts.GetOrAskForADescription(ctx)
	messaging.Tags = prompts.GetOrAskForTags(ctx)

	messaging.Local = prompts.GetOrAskForLocal(ctx)
	messaging.Regex = prompts.GetMatchRegex(ctx)
	messaging.Match = GetOrRequireAChannelMatch(ctx)
	messaging.MQTT = GetMQTT(ctx)
	messaging.WebSocket = GetWebSocket(ctx)

	return messaging, nil
}
