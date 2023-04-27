package common

// TODO: Add dreamland connection
// TODO: Move to specs
const (
	DefaultNetwork    = "Taubyte's Sandbox Network"
	DeprecatedNetwork = "Sandbox Network [Deprecated]"
	CustomNetwork     = "Custom"
	PythonTestNetwork = "Test"
)

var (
	SelectedNetwork = DefaultNetwork
	NetworkTypes    = []string{DefaultNetwork, DeprecatedNetwork, CustomNetwork}

	DefaultAuthUrl    = "https://auth.tau.sandbox.taubyte.com"
	DefaultPatrickUrl = "https://patrick.tau.sandbox.taubyte.com"
	DefaultSeerUrl    = "https://seer.tau.sandbox.taubyte.com"

	DeprecatedAuthUrl    = "https://auth.taubyte.com"
	DeprecatedPatrickUrl = "https://patrick.taubyte.com"
	DeprecatedSeerUrl    = "https://seer.taubyte.com"
)
