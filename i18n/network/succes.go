package network

import (
	"github.com/pterm/pterm"
	"github.com/taubyte/tau/common"
)

func Success(network, fqdn string) {
	if network == common.CustomNetwork {
		pterm.Success.Printfln("Connected to custom network fqdn: %s", pterm.FgCyan.Sprintf(fqdn))
	} else {
		pterm.Success.Printfln("Connected to %s", pterm.FgCyan.Sprintf(network))
	}

}
