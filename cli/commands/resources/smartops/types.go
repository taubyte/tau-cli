package smartops

import (
	"github.com/taubyte/tau-cli/cli/common"
)

type link struct{}

// New is called in tau/cli/new.go to attach the relative commands
// to their parents, i.e `new` => `new smartops`
func New() common.Basic {
	return link{}
}

func (link) Clone() common.Command {
	return common.NotImplemented
}

func (link) Select() common.Command {
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
