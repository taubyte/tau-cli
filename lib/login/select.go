package loginLib

import (
	"github.com/taubyte/tau/env"
	loginI18n "github.com/taubyte/tau/i18n/login"
	"github.com/taubyte/tau/singletons/config"
	"github.com/urfave/cli/v2"
)

/*
if setDefault is true it will remove current default and set the
newly selected profile as the default
*/
func Select(ctx *cli.Context, name string, setDefault bool) error {
	if setDefault == true {
		configProfiles := config.Profiles()
		profiles := configProfiles.List(true)
		for profileName, profile := range profiles {
			if profileName == name {
				profile.Default = true
				err := configProfiles.Set(profileName, profile)
				if err != nil {
					return loginI18n.SettingDefaultFailed(err)
				}
				continue
			}

			if profile.Default == true {
				profile.Default = false

				err := configProfiles.Set(profileName, profile)
				if err != nil {
					return loginI18n.RemovingDefaultFailed(err)
				}
			}
		}
	}

	profile, err := config.Profiles().Get(name)
	if err != nil {
		return err
	}

	env.SetSelectedNetwork(ctx, profile.Network)
	return env.SetSelectedUser(ctx, name)
}
