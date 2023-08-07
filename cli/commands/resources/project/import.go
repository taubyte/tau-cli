package project

import (
	"fmt"

	"github.com/taubyte/tau-cli/cli/common"
	"github.com/urfave/cli/v2"
)

func (link) Import() common.Command {
	return common.Create(
		&cli.Command{
			Action: _import,
		},
	)
}

func _import(ctx *cli.Context) error {
	fmt.Println("HELLO WORLD")
	return nil
}
