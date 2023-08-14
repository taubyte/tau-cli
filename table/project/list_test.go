package projectTable_test

import (
	"fmt"

	projectTable "github.com/taubyte/tau-cli/table/project"
	client "github.com/taubyte/tau/clients/http/auth"
)

func ExampleList() {
	project := []*client.Project{
		{
			Id:   "QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH",
			Name: "someProject1",
		},
		{
			Id:   "QmbUIDhRosp5BaXDASEWSCtpkQCgQCPdRVhnxjiSHfXdC0",
			Name: "someProject2",
		},
	}

	projectTable.List(project, func(project *client.Project) string {
		return fmt.Sprintf("This is a description of `%s` it does cool stuff I promise", project.Name)
	})

	// Output:
	// ┌─────────────────┬──────────────┬──────────────────────────────────────────┐
	// │ ID              │ NAME         │ DESCRIPTION                              │
	// ├─────────────────┼──────────────┼──────────────────────────────────────────┤
	// │ QmbAA8...HfXdWH │ someProject1 │ This is a description of `someProject1`  │
	// │                 │              │ it does cool stuff I promise             │
	// ├─────────────────┼──────────────┼──────────────────────────────────────────┤
	// │ QmbUID...HfXdC0 │ someProject2 │ This is a description of `someProject2`  │
	// │                 │              │ it does cool stuff I promise             │
	// └─────────────────┴──────────────┴──────────────────────────────────────────┘
}
