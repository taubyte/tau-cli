package projectPrompts

const (
	ProjectName       = "Project Name:"
	ProjectVisibility = "Visibility:"
	SelectAProject    = "Select a Project:"

	CreateThisProject       = "Create this project?"
	Public                  = "public"
	Private                 = "private"
	NoProjectsFound         = "no projects found"
	NoProjectsWithNameFound = "no projects with name `%s` found"
)

var (
	VisibilityOptions = []string{Public, Private}
)
