package projectI18n

import (
	"github.com/taubyte/tau-cli/i18n/printer"
)

func success(prefix, name string) {
	printer.SuccessWithName("%s project: %s", prefix, name)
}

func DeselectedProject(name string) {
	success("Deselected", name)
}

func SelectedProject(name string) {
	success("Selected", name)
}

func CreatedProject(name string) {
	success("Created", name)
}

func PushedProject(name string) {
	success("Pushed", name)
}

func PulledProject(name string) {
	success("Pulled", name)
}

func CheckedOutProject(name, branch string) {
	printer.SuccessWithName("Checked out branch `%s` on project `%s`", branch, name)
}
