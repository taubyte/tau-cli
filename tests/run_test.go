//go:build !cover

package tests

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/taubyte/tau/constants"
)

func takeCover() func() {
	return func() {}
}

func (r *roadRunner) Run(args ...string) (error, int, string, string) {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting cwd failed with: %s", err), 1, "", ""
	}

	_cmd := exec.Command(path.Join(cwd, "tau"), args...)
	_cmd.Dir, _ = filepath.Abs(r.dir)
	r.env[constants.TauConfigFileNameEnvVarName] = r.configFile
	r.env[constants.TauSessionLocationEnvVarName] = r.sessionFile
	if r.authUrl != "" {
		// All the tests run with localAuthClient tag, but if the mock tag is false
		// the url is still set to the auth.taubyte.com url
		r.env[constants.AuthURLEnvVarName] = r.authUrl
	}

	_cmd.Env = os.Environ()

	for k, v := range r.env {
		_cmd.Env = append(_cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	// Capture command output
	var out bytes.Buffer
	var errOut bytes.Buffer
	_cmd.Stdout = &out
	_cmd.Stderr = &errOut

	// Start the command
	err = _cmd.Start()
	if err != nil {
		return fmt.Errorf("Command failed to start %s", err.Error()), 1, "", ""
	}

	// Kill the command after the timeout
	done := make(chan bool)
	go func() {
		select {
		case <-time.After(r.waitTime):
			err = _cmd.Process.Kill()
			if err != nil {
				panic(err)
			}
		case <-done:
			return
		}
	}()

	// Wait for the command to finish
	err = _cmd.Wait()
	if err != nil {
		exiterr := err.(*exec.ExitError)
		status := exiterr.Sys().(syscall.WaitStatus)
		return err, status.ExitStatus(), out.String(), errOut.String()
	}
	return nil, 0, out.String(), errOut.String()
}
