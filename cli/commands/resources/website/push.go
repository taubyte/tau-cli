package website

import (
	"github.com/taubyte/tau/cli/common"
)

func (l link) Push() common.Command {
	return l.cmd.PushCmd()
}
