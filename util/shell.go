package util

import (
	"os"
	"os/exec"
)

func ShellCommand(cmd string, args ...string) string {
	fullShellCommand := exec.Command(cmd, args...)
	stdout, err := fullShellCommand.Output()

	if err != nil {
		panic(err.Error())
	}

	return string(stdout)
}

func ShellCommandWithInput(inputFileName string, cmd string, args ...string) string {
	processStdin, _ := os.Open(inputFileName)

	fullShellCommand := exec.Command(cmd, args...)
	fullShellCommand.Stdin = processStdin
	stdout, err := fullShellCommand.Output()

	if err != nil {
		panic(err.Error())
	}

	return string(stdout)
}
