package library

import (
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/urfave/cli/v2"
)

func (link) Import() common.Command {
	return common.Create(
		&cli.Command{},
	)
}
