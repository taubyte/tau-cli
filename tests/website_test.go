package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	commonTest "github.com/taubyte/tau/common/test"
	"github.com/taubyte/tau/constants"
	websitePrompts "github.com/taubyte/tau/prompts/website"
)

func TestWebsiteAll(t *testing.T) {
	runTests(t, createWebsiteMonkey(), true)
}

func createWebsiteMonkey() *testSpider {
	// Define shared variables
	command := "website"
	profileName := "test"
	projectName := "test_project"
	testName := "someWebsite"

	testDomain := "test_domain_1"
	testDomainFqdn := "hal.computers.com"
	network := "Sandbox Network [Deprecated]"

	// Create a basic resource of name
	basicNew := func(name string) []string {
		return []string{
			"new", "-y", command,
			"--name", name,
			"--description", "some website description",
			"--tags", "tag1, tag2,   tag3",
			"--no-generate-repository",
			"--paths", "/",
			"--repository-name", "tb_website_reactdemo",
			"--no-clone",
			"--branch", "master",
			"--provider", "github",
			"--domains", testDomain,
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
				"--description", "some website description",
				"--tags", "tag1, tag2,   tag3",
				"--generate-repository",
				"--private",
				"--template", "html",
				"--branch", "master",
				"--paths", "/",
				"--domains", testDomain,
				"--provider", "github",
				"--no-embed-token",
			},
			writeFilesInDir: specialWriteFilesInDir(testDomainFqdn),
			preRun: [][]string{
				basicNewDomain(testDomain, testDomainFqdn),
			},
			cleanUp: func() error {
				// delete https://github.com/taubyte-test/tb_empty_template
				url := fmt.Sprintf("https://api.github.com/repos/%s/tb_website_someWebsite", commonTest.GitUser)
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
			name: "New from current repository",
			args: []string{
				"query", command, testName,
			},
			// TODO confirm values
			preRun: [][]string{
				basicNewDomain(testDomain, testDomainFqdn),
				basicNew(testName),
			},
			writeFilesInDir: specialWriteFilesInDir(testDomainFqdn),
		},
		{
			name: "edit basic",
			args: []string{
				"query", command, testName,
			},
			// TODO confirm values
			preRun: [][]string{
				basicNewDomain(testDomain, testDomainFqdn),
				basicNewDomain("test_domain2", "test.computers.com"),
				basicNew(testName),
			},
			writeFilesInDir: specialWriteFilesInDir(testDomainFqdn),
			children: []testMonkey{
				{
					name: "edit",
					args: []string{
						"edit", "-y", command, testName,
						"--description", "some new website description",
						"--tags", "tag1, tag2,   tag4",
						"--paths", "/",
						// "--repository-name", "tb_website_reactdemo",
						"--no-clone",
						"--branch", "master",
						"--domains", "test_domain2",
					},
					// TODO confirm values
				},
			},
		},
		{
			name: "delete basic",
			args: []string{
				"query", command, testName,
			},
			exitCode: 1,
			errOut:   []string{fmt.Sprintf(websitePrompts.NotFound, testName)},
			preRun: [][]string{
				basicNewDomain(testDomain, testDomainFqdn),
				basicNew(testName),
			},
			writeFilesInDir: specialWriteFilesInDir(testDomainFqdn),
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
				basicNewDomain(testDomain, testDomainFqdn),
				basicNew(testName + "1"),
				basicNew(testName + "2"),
				basicNew(testName + "3"),
				{"delete", "-y", command, "--name", testName + "3"},
				basicNew(testName + "4"),
				basicNew(testName + "5"),
			},
			writeFilesInDir: specialWriteFilesInDir(testDomainFqdn),
		},
	}

	return &testSpider{projectName, tests, beforeEach, getConfigString, "website"}
}
