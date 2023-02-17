package templates

func Get() *templates {
	getOrCreateTemplates()

	return _templates
}

type templateYaml struct {
	// parameters must be exported for the yaml parser
	Name        string
	Description string
	Icon        string
	URL         string
}
