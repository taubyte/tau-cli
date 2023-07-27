package messagingPrompts

import (
	messagingFlags "github.com/taubyte/tau-cli/flags/messaging"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/urfave/cli/v2"
)

func GetOrRequireAChannelMatch(ctx *cli.Context, prev ...string) string {
	return prompts.GetOrRequireAMatch(ctx, ChannelMatch, prev...)
}

func GetMQTT(ctx *cli.Context, prev ...bool) bool {
	return prompts.GetOrAskForBool(ctx, messagingFlags.MQTT.Name, MQTTPrompt, prev...)
}

func GetWebSocket(ctx *cli.Context, prev ...bool) bool {
	return prompts.GetOrAskForBool(ctx, messagingFlags.WebSocket.Name, WebSocketPrompt, prev...)
}
