package library

import (
	"github.com/taubyte/tau/cli/common"
)

func (l link) Clone() common.Command {
	return l.cmd.CloneCmd()
}
