package functionPrompts

import (
	"os"
	"path"

	"github.com/taubyte/go-project-schema/functions"
	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau-cli/flags"
	functionFlags "github.com/taubyte/tau-cli/flags/function"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/taubyte/tau-cli/singletons/templates"
	"github.com/urfave/cli/v2"
)

// Should only be used in new, will overwrite values.
func checkTemplate(ctx *cli.Context, function *structureSpec.Function) (templateURL string, err error) {
	if !ctx.IsSet(flags.Template.Name) && !prompts.GetUseACodeTemplate(ctx) {
		return
	}

	language, err := prompts.GetOrSelectLanguage(ctx)
	if err != nil {
		return
	}

	codeTemplateMap, err := templates.Get().Function(language)
	if err != nil {
		return
	}

	templateURL, err = prompts.SelectATemplate(ctx, codeTemplateMap)
	if err != nil {
		return
	}

	file, err := os.ReadFile(path.Join(templateURL, "config.yaml"))
	if err != nil {
		return
	}

	getter, err := functions.Yaml(function.Name, "", file)
	if err != nil {
		return
	}

	_function, err := getter.Struct()
	if err != nil {
		return
	}

	// Overwrite new function
	*function = *_function

	return
}

func editHttp(ctx *cli.Context, function *structureSpec.Function) (err error) {
	function.Domains, err = prompts.GetOrSelectDomainsWithFQDN(ctx, function.Domains...)
	if err != nil {
		return
	}

	function.Method, err = GetHttpMethod(ctx, function.Method)
	if err != nil {
		return
	}

	function.Paths = prompts.RequiredPaths(ctx, function.Paths...)
	return
}

func editP2P(ctx *cli.Context, function *structureSpec.Function) (err error) {
	function.Protocol, err = prompts.SelectAServiceWithProtocol(ctx, functionFlags.Protocol.Name, ProtocolPrompt, function.Protocol)
	if err != nil {
		return
	}

	function.Command = GetOrRequireACommand(ctx, function.Command)

	function.Local = prompts.GetOrAskForLocal(ctx, function.Local)
	return
}

func editPubSub(ctx *cli.Context, function *structureSpec.Function) (err error) {
	function.Channel = GetOrRequireAChannel(ctx, function.Channel)
	function.Local = prompts.GetOrAskForLocal(ctx, function.Local)
	return
}
