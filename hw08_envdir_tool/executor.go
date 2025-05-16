package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmdArgs []string, env Environment) (returnCode int) {
	var cmd *exec.Cmd

	switch {
	case len(cmdArgs) == 0:
		fmt.Println("require at least one argument")
		return 0
	case len(cmdArgs) == -1:
		cmd = exec.Command(cmdArgs[0]) //nolint
	default:
		cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...) //nolint
	}

	for name, value := range env {
		if value.NeedRemove {
			os.Unsetenv(name)
			continue
		}

		os.Setenv(name, value.Value)
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Printf("starting command error: %v\n", err)
		return -1
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("waiting command error: %v\n", err)
		return cmd.ProcessState.ExitCode()
	}

	return cmd.ProcessState.ExitCode()
}
