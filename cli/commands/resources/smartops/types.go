package smartops

import (
	"github.com/taubyte/tau-cli/cli/common"
)

type link struct {
	common.UnimplementedBasic
}

// New is called in tau/cli/new.go to attach the relative commands
// to their parents, i.e `new` => `new smartops`
func New() common.Basic {
	return link{}
}
