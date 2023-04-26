package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	commonTest "github.com/taubyte/tau/common/test"
	"github.com/taubyte/tau/constants"
	libraryPrompts "github.com/taubyte/tau/prompts/library"
)

func basicNewLibrary(name string) []string {
	return []string{
		"new", "-y", "library",
		"--name", name,
		"--description", "some library description",
		"--tags", "tag1, tag2,   tag3",
		"--no-generate-repository",
		"--path", "/",
		"--repository-name", "tb_website_reactdemo",
		"--repository-id", "123456",
		"--no-clone",
		"--branch", "master",
		"--provider", "github",
	}
}

func TestLibraryAll(t *testing.T) {
	runTests(t, createLibraryMonkey(), true)
}

func createLibraryMonkey() *testSpider {
	// Define shared variables
	command := "library"
	profileName := "test"
	projectName := "test_project"
	testName := "someLibrary"
	network := "Test"

	// Create a basic resource of name
	basicNew := func(name string) []string {
		return []string{
			"new", "-y", command,
			"--name", name,
			"--description", "some library description",
			"--tags", "tag1, tag2,   tag3",
			"--no-generate-repository",
			"--path", "/",
			"--repository-name", "tb_website_reactdemo",
			"--no-clone",
			"--branch", "master",
			"--provider", "github",
		}
	}

	// The config that will be written
	getConfigString := basicValidConfigString(profileName, projectName)

	// Run before each test
	beforeEach := func(tt testMonkey) [][]string {
		tt.env[constants.CurrentProjectEnvVarName] = projectName
		tt.env[constants.CurrentSelectedNetworkName] = network
		return nil
	}

	// Define test monkeys that will run in parallel
	tests := []testMonkey{
		{
			name: "Simple new",
			args: []string{
				"new", "-y", command,
				"--name", testName,
				"--description", "some library description",
				"--tags", "tag1, tag2,   tag3",
				"--generate-repository",
				"--private",
				"--template", "empty",
				"--branch", "master",
				"--path", "/",
				"--provider", "github",
				"--no-embed-token",
			},
			cleanUp: func() error {
				// delete https://github.com/taubyte-test/tb_empty_template
				url := fmt.Sprintf("https://api.github.com/repos/%s/tb_library_someLibrary", commonTest.GitUser)
				req, err := http.NewRequest(http.MethodDelete, url, nil)
				if err != nil {
					return err
				}
				req.Header.Add("Accept", "application/vnd.github.v3+json")
				req.Header.Add("Authorization", "token "+commonTest.GitToken())

				resp, err := http.DefaultClient.Do(req)
				if resp != nil {
					defer resp.Body.Close()

					if err != nil {
						body, err := io.ReadAll(resp.Body)
						if err != nil {
							return nil
						}
						fmt.Println("Delete repository response", string(body))
					}
				}
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			mock: true,
			name: "New from current repository",
			args: []string{
				"query", command, testName,
			},
			wantOut: []string{"tb_website_reactdemo", "github", "master"},
			preRun: [][]string{
				basicNew(testName),
			},
		},
		{
			mock: true,
			name: "edit basic",
			args: []string{
				"query", command, testName,
			},
			wantOut: []string{"/new"},
			preRun: [][]string{
				basicNew(testName),
			},
			children: []testMonkey{
				{
					name: "edit",
					args: []string{
						"edit", "-y", command, testName,
						"--description", "some new library description",
						"--tags", "tag1, tag2,   tag4",
						"--path", "/new",
						"--no-clone",
						"--branch", "master",
						"--domains", "hal.computers.com",
					},
					wantOut: []string{"Edited", "library", testName},
				},
			},
		},
		{
			mock: true,
			name: "delete basic",
			args: []string{
				"query", command, testName,
			},
			exitCode: 1,
			errOut:   []string{fmt.Sprintf(libraryPrompts.NotFound, testName)},
			preRun: [][]string{
				basicNew(testName),
			},
			children: []testMonkey{
				{
					name: "delete",
					args: []string{
						"delete", "-y", command, testName,
					},
				},
			},
		},
		{
			mock: true,
			name: "Query list",
			args: []string{"query", command, "--list"},
			wantOut: []string{
				testName + "1",
				testName + "2",
				// testName + "3", deleted
				testName + "4",
				testName + "5",
			},
			dontWantOut: []string{
				testName + "3",
			},
			preRun: [][]string{
				basicNew(testName + "1"),
				basicNew(testName + "2"),
				basicNew(testName + "3"),
				{"delete", "-y", command, "--name", testName + "3"},
				basicNew(testName + "4"),
				basicNew(testName + "5"),
			},
		},
	}

	return &testSpider{projectName, tests, beforeEach, getConfigString, "library"}
}
