//go:build !no_rebuild

package tests

// Always rebuild between `go test ...` command executions
func buildTau() error {
	return internalBuildTau()
}
