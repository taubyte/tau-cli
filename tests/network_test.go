package tests

import (
	"testing"

	"github.com/pterm/pterm"
	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/constants"
)

func TestNetworkAll(t *testing.T) {
	runTests(t, createNetworkMonkey(), true)
}

func createNetworkMonkey() *testSpider {
	// Define shared variables
	profileName := "test"
	projectName := "test_project"
	customNetwork := common.CustomNetwork
	deprecatedNetwork := common.DeprecatedNetwork
	defaultNetwork := common.DefaultNetwork
	fqdn := "aron.lol"

	// The config that will be written
	getConfigString := basicGetConfigString(profileName, projectName)

	// Run before each test
	beforeEach := func(tt testMonkey) [][]string {
		tt.env[constants.CurrentSelectedNetworkName] = ""
		return nil
	}

	tests := []testMonkey{
		{
			name: "Select default network",
			args: []string{"select", "network", "--default"},
			wantOut: []string{
				pterm.Success.Sprintf("Connected to %s", pterm.FgCyan.Sprintf(defaultNetwork)),
			},
			evaluateSession: expectedSelectedNetwork(defaultNetwork),
		},
		{
			name: "Select deprecated network",
			args: []string{"select", "network", "--deprecated"},
			wantOut: []string{
				pterm.Success.Sprintf("Connected to %s", pterm.FgCyan.Sprintf(deprecatedNetwork)),
			},
			evaluateSession: expectedSelectedNetwork(deprecatedNetwork),
		},
		{
			name: "Select custom network with FQDN",
			args: []string{"select", "network", "--fqdn", fqdn},
			wantOut: []string{
				pterm.Success.Sprintf("Connected to custom network fqdn: %s", pterm.FgCyan.Sprintf(fqdn)),
			},
			evaluateSession: expectedSelectedCustomNetwork(customNetwork, fqdn),
		},
		{
			name:            "Select login with network saved",
			args:            []string{"login", "--name", profileName},
			evaluateSession: expectedSelectedNetwork(deprecatedNetwork),
		},
	}

	return &testSpider{projectName, tests, beforeEach, getConfigString, "network"}
}
