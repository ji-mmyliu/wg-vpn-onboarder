package util

import (
	"os/exec"
)

func RunShellCommand(cmd string, args ...string) string {
	fullShellCommand := exec.Command(cmd, args...)
	stdout, err := fullShellCommand.Output()

	if err != nil {
		panic(err.Error())
	}

	return string(stdout)
}
