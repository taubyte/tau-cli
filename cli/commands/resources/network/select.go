package network

import (
	cliCommon "github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/env"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/prompts"
	domainPrompts "github.com/taubyte/tau-cli/prompts/domain"
	"github.com/taubyte/tau-cli/singletons/config"

	networkFlags "github.com/taubyte/tau-cli/flags/network"
	networkI18n "github.com/taubyte/tau-cli/i18n/network"
	"github.com/urfave/cli/v2"
)

func (link) Select() cliCommon.Command {
	return cliCommon.Create(
		&cli.Command{
			Action: _select,
			Flags:  []cli.Flag{networkFlags.FQDN, networkFlags.Default, networkFlags.Deprecated},
		},
	)
}

// TODO: maybe save the old custom url when switching and prompt to use same as saved before
func _select(ctx *cli.Context) error {
	var selectedNetwork, fqdn string

	// Setting string flag with value counts as two
	if ctx.NumFlags() > 1 {
		if !ctx.IsSet(networkFlags.FQDN.Name) {
			return networkI18n.FlagError()
		}
	}

	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return err
	}

	if ctx.Bool(networkFlags.Default.Name) {
		selectedNetwork = common.DefaultNetwork
		profile.FQDN = ""
	} else if ctx.Bool(networkFlags.Deprecated.Name) {
		selectedNetwork = common.DeprecatedNetwork
		profile.FQDN = ""
	} else if fqdn = ctx.String(networkFlags.FQDN.Name); fqdn != "" {
		selectedNetwork = common.CustomNetwork
		env.SetCustomNetworkUrl(ctx, fqdn)
		profile.FQDN = fqdn
	} else {
		selectedNetwork = prompts.GetOrAskForSelection(ctx, "", prompts.NetworkPrompts, common.NetworkTypes, common.SelectedNetwork)
		if selectedNetwork == common.CustomNetwork {
			fqdn = domainPrompts.GetOrRequireAnFQDN(ctx, "")
			env.SetCustomNetworkUrl(ctx, fqdn)
			profile.FQDN = fqdn
		} else {
			profile.FQDN = ""
		}

	}

	profile.Network = selectedNetwork

	config.Profiles().Set(profile.Name(), profile)
	if err := env.SetSelectedNetwork(ctx, selectedNetwork); err != nil {
		return err
	}

	networkI18n.Success(selectedNetwork, fqdn)

	return nil
}
