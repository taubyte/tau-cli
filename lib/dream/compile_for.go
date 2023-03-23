package dreamLib

type CompileForDFunc struct {
	ProjectId     string
	ApplicationId string
	ResourceId    string
	Branch        string
	Call          string
	Path          string
}

func (c *CompileForDFunc) Execute() error {
	args := []string{
		"inject", "compileFor",
		"--project-id", c.ProjectId,
		"--resource-id", c.ResourceId,
		"--branch", c.Branch,
		"--call", c.Call,
		"--path", c.Path,
	}

	if len(c.ApplicationId) > 0 {
		args = append(args, []string{"--application-id", c.ApplicationId}...)
	}

	return Execute(args...)
}

type CompileForRepository struct {
	ProjectId     string
	ApplicationId string
	ResourceId    string
	Branch        string
	Path          string
}

func (c *CompileForRepository) Execute() error {
	args := []string{
		"inject", "compileFor",
		"--project-id", c.ProjectId,
		"--resource-id", c.ResourceId,
		"--branch", c.Branch,
		"--path", c.Path,
	}

	if len(c.ApplicationId) > 0 {
		args = append(args, []string{"--application-id", c.ApplicationId}...)
	}

	return Execute(args...)
}
