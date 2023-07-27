package env

import (
	"fmt"

	"github.com/taubyte/tau-cli/flags"
	"github.com/urfave/cli/v2"
)

func justDisplayExport(c *cli.Context, key, value string) bool {
	if c.Bool(flags.Env.Name) == true {
		fmt.Printf("export %s=%s\n", key, value)
		return true
	}

	return false
}
