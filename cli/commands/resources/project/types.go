package project

import (
	"github.com/taubyte/tau-cli/cli/common"
)

type link struct {
}

func New() common.Basic {
	return link{}
}

func (link) Edit() common.Command {
	return common.NotImplemented
}

func (link) Delete() common.Command {
	return common.NotImplemented
}
