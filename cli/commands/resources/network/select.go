package network

import (
	"fmt"

	cliCommon "github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/env"
	loginLib "github.com/taubyte/tau-cli/lib/login"
	"github.com/taubyte/tau-cli/prompts"
	"github.com/taubyte/tau-cli/singletons/config"
	"github.com/taubyte/tau-cli/singletons/dreamland"
	"github.com/taubyte/tau-cli/validate"
	slices "github.com/taubyte/utils/slices/string"

	networkFlags "github.com/taubyte/tau-cli/flags/network"
	networkI18n "github.com/taubyte/tau-cli/i18n/network"
	"github.com/urfave/cli/v2"
)

func (link) Select() cliCommon.Command {
	return cliCommon.Create(
		&cli.Command{
			Action: _select,
			Flags:  []cli.Flag{networkFlags.FQDN, networkFlags.Universe},
		},
	)
}

func _select(ctx *cli.Context) error {
	// Setting string flag with value counts as two
	if ctx.NumFlags() > 2 {
		return networkI18n.FlagError()
	}

	profile, err := loginLib.GetSelectedProfile()
	if err != nil {
		return err
	}

	switch {
	case ctx.IsSet(networkFlags.FQDN.Name):
		profile.NetworkType = common.RemoteNetwork
		profile.Network = ctx.String(networkFlags.FQDN.Name)

		if err := validate.SeerFQDN(ctx.Context, profile.Network); err != nil {
			return err
		}

		if !slices.Contains(profile.History, profile.Network) {
			profile.History = append(profile.History, profile.Network)
		}

	case ctx.IsSet(networkFlags.Universe.Name):
		dreamClient, err := dreamland.Client(ctx.Context)
		if err != nil {
			return fmt.Errorf("creating dreamland client failed with: %w", err)
		}

		universes, err := dreamClient.Status()
		if err != nil {
			return fmt.Errorf("calling dreamland status failed with: %w", err)
		}

		universeName := ctx.String(networkFlags.Universe.Name)
		_, ok := universes[universeName]
		if !ok {
			return fmt.Errorf("universe `%s` was not found", universeName)
		}

		profile.NetworkType = common.DreamlandNetwork
		profile.Network = universeName
	default:
		dreamClient, err := dreamland.Client(ctx.Context)
		if err != nil {
			return fmt.Errorf("creating dreamland client failed with: %w", err)
		}

		networkSelections := []string{common.RemoteNetwork}
		if _, err := dreamClient.Status(); err == nil {
			networkSelections = append(networkSelections, common.DreamlandNetwork)
		}

		networkSelections = append(networkSelections, profile.History...)

		prev := []string{}
		if len(profile.NetworkType) > 0 {
			prev = append(prev, profile.NetworkType)
		}

		network := prompts.GetOrAskForSelection(ctx, "Network", prompts.NetworkPrompts, networkSelections, prev...)
		if network == common.RemoteNetwork {
			profile.NetworkType = common.RemoteNetwork
			profile.Network = prompts.GetOrRequireAString(ctx, "", prompts.FQDN, validate.FQDNValidator, profile.Network)
			if err := validate.SeerFQDN(ctx.Context, profile.Network); err != nil {
				return err
			}

			if !slices.Contains(profile.History, profile.Network) {
				profile.History = append(profile.History, profile.Network)
			}

		} else if network == common.DreamlandNetwork {
			universes, err := dreamClient.Status()
			if err != nil {
				return fmt.Errorf("calling dreamland status failed with: %w", err)
			}

			universeNames := make([]string, 0, len(universes))
			for name := range universes {
				universeNames = append(universeNames, name)
			}

			profile.Network, err = prompts.SelectInterface(universeNames, prompts.Universe, "")
			if err != nil {
				return fmt.Errorf("universe selection failed with: %w", err)
			}
		} else {
			profile.NetworkType = common.RemoteNetwork
			profile.Network = network
		}
	}

	config.Profiles().Set(profile.Name(), profile)
	if err := env.SetSelectedNetwork(ctx, profile.NetworkType); err != nil {
		return err
	}

	if err := env.SetNetworkUrl(ctx, profile.Network); err != nil {
		return err
	}

	networkI18n.Success(profile.Network)

	return nil
}
