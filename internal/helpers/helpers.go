package helpers

import (
	"os"
	"os/exec"
)

// ExecuteCmd is a utility function to execute shell commands.
func ExecuteCmd(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
