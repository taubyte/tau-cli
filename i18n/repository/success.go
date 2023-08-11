package repositoryI18n

import "github.com/taubyte/tau-cli/i18n/printer"

func Imported(name, network string) {
	printer.SuccessWithNameOnNetwork("Imported %s: %s on network %s", "repository", name, network)
}
