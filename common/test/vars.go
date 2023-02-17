package internal

import "os"

type Repository struct {
	ID   int
	Name string
	URL  string
}

func GitToken() string {
	token := os.Getenv("TAU_TEST_GIT_TOKEN")

	if token == "" {
		panic("TAU_TEST_GIT_TOKEN not set")
	}

	return token
}

var (
	GitUser     = "taubyte-test"
	Branch      = "master"
	ProjectName = "testproject"

	ConfigRepo Repository = Repository{
		ID:   485473636,
		Name: "tb_testproject",
		URL:  "https://github.com/taubyte-test/tb_testproject",
	}

	CodeRepo Repository = Repository{
		ID:   485473661,
		Name: "tb_code_testproject",
		URL:  "https://github.com/taubyte-test/tb_code_testproject",
	}
)
