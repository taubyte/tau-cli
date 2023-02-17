package printer

import "github.com/pterm/pterm"

func SuccessWithName(format, prefix, name string) {
	pterm.Success.Printfln(format, prefix, pterm.FgCyan.Sprintf(name))
}
