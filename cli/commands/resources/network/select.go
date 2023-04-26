package network

import (
	cliCommon "github.com/taubyte/tau/cli/common"
	"github.com/taubyte/tau/common"
	"github.com/taubyte/tau/env"
	"github.com/taubyte/tau/prompts"
	domainPrompts "github.com/taubyte/tau/prompts/domain"
	"github.com/taubyte/tau/singletons/config"

	networkFlags "github.com/taubyte/tau/flags/network"
	networkI18n "github.com/taubyte/tau/i18n/network"
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

	selectedProfile, _ := env.GetSelectedUser()
	profile, err := config.Profiles().Get(selectedProfile)
	if err != nil {
		return err
	}

	if ctx.Bool(networkFlags.Default.Name) {
		selectedNetwork = common.DefaultNetwork
		profile.FQDN = ""
	} else if ctx.Bool(networkFlags.Deprecated.Name) {
		selectedNetwork = common.DeprecatedNetwork
		profile.FQDN = ""
	} else if ctx.String(networkFlags.FQDN.Name) != "" {
		selectedNetwork = common.CustomNetwork
		fqdn = ctx.String(networkFlags.FQDN.Name)
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

	config.Profiles().Set(selectedProfile, profile)
	err = env.SetSelectedNetwork(ctx, selectedNetwork)
	if err != nil {
		return err
	}

	networkI18n.Success(selectedNetwork, fqdn)

	return nil
}
