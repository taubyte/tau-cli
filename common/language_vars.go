package common

import "github.com/taubyte/go-specs/builders/wasm"

func GetLanguages() (languages []string) {
	for lang := range wasm.SupportedLanguages() {
		languages = append(languages, string(lang))
	}

	return
}
