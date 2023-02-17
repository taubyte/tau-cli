package session

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-ps"
)

func parentId(pid int) (int, error) {
	process, err := ps.FindProcess(pid)
	if err != nil {
		return 0, err
	}

	// This is for `go run .`
	if process.Executable() == "go" {
		return process.PPid(), nil
	}

	return process.Pid(), nil
}

func discoverOrCreateConfigFileLoc() (string, error) {
	grandPid, err := parentId(os.Getppid())
	if err != nil {
		return "", err
	}

	processDir, found := nearestProcessDirectory(grandPid)
	if found == false {
		processDir, err = createProcessDirectory(grandPid)
		if err != nil {
			return "", err
		}
	}

	return processDir, nil
}

/*
Nearest process directory will climb up the process tree until it finds a
directory or the pid reaches 1
*/
func nearestProcessDirectory(pid int) (processDir string, found bool) {
	processDir = directoryFromPid(pid)

	_, err := os.Stat(processDir)
	if err != nil {
		process, err := ps.FindProcess(pid)
		if err != nil {
			return
		}

		ppid := process.PPid()
		if ppid == 1 {
			return
		}

		processDir = directoryFromPid(ppid)

		_, err = os.Stat(processDir)
		if err != nil {
			return nearestProcessDirectory(ppid)
		}
	}

	return processDir, true
}

func directoryFromPid(pid int) string {
	return path.Join(os.TempDir(), fmt.Sprintf("%s-%d", sessionDirPrefix, pid))
}

func createProcessDirectory(pid int) (string, error) {
	processDir := directoryFromPid(pid)

	err := os.Mkdir(processDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return processDir, nil
}
