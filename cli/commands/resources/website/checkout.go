package website

import (
	"github.com/taubyte/tau-cli/cli/common"
)

func (l link) Checkout() common.Command {
	return l.cmd.CheckoutCmd()
}
