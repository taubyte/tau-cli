package network

import (
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/urfave/cli/v2"
)

func (link) Base() (*cli.Command, []common.Option) {
	return common.Base(
		&cli.Command{
			Name: "network",
		},
	)
}
