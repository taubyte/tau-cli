package library

import (
	"github.com/taubyte/tau-cli/cli/common"
)

func (l link) Pull() common.Command {
	return l.cmd.PullCmd()
}
