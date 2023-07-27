package website

import (
	"github.com/taubyte/tau-cli/cli/common"
)

func (l link) Clone() common.Command {
	return l.cmd.CloneCmd()
}
