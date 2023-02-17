package website

import (
	"github.com/taubyte/tau/cli/common"
)

func (l link) Pull() common.Command {
	return l.cmd.PullCmd()
}
