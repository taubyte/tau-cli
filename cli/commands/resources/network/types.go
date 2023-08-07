package network

import (
	"github.com/taubyte/tau-cli/cli/common"
)

type link struct {
	common.UnimplementedBasic
}

func New() common.Basic {
	return link{}
}

func (link) New() common.Command {
	return nil
}
