package functionPrompts

import (
	"github.com/taubyte/tau/common"
	functionFlags "github.com/taubyte/tau/flags/function"
	"github.com/taubyte/tau/prompts"
	"github.com/urfave/cli/v2"
)

func GetFunctionType(ctx *cli.Context, prev ...string) (string, error) {
	return prompts.SelectInterfaceField(ctx,
		common.FunctionTypes,
		functionFlags.Type.Name,
		TypePrompt,
		prev...,
	)
}

func GetHttpMethod(ctx *cli.Context, prev ...string) (string, error) {
	return prompts.SelectInterfaceField(ctx,
		common.HTTPMethodTypes,
		functionFlags.Method.Name,
		MethodPrompt,
		prev...,
	)
}

func GetOrRequireACommand(ctx *cli.Context, prev ...string) string {
	return prompts.GetOrRequireAString(ctx,
		functionFlags.Command.Name,
		CommandPrompt,
		nil,
		prev...,
	)
}

func GetOrRequireAChannel(ctx *cli.Context, prev ...string) string {
	return prompts.GetOrRequireAString(ctx,
		functionFlags.Channel.Name,
		ChannelPrompt,
		nil,
		prev...,
	)
}
