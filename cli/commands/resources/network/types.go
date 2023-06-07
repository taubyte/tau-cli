package network

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

func (link) Edit() common.Command {
	return common.NotImplemented
}

func (link) List() common.Command {
	return common.NotImplemented
}

func (link) Delete() common.Command {
	return common.NotImplemented

}

func (link) New() common.Command {
	return nil

}

func (link) Query() common.Command {
	return common.NotImplemented

}
