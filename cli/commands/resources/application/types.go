package application

import (
	"github.com/taubyte/tau/cli/common"
)

type link struct{}

func New() common.Basic {
	return link{}
}

func (link) Clone() common.Command {
	return common.NotImplemented
}

func (link) Push() common.Command {
	return common.NotImplemented
}

func (link) Pull() common.Command {
	return common.NotImplemented
}

func (link) Checkout() common.Command {
	return common.NotImplemented
}
