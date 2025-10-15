package cmd

import (
	"os"
	"os/exec"
)

// RunShellCommand executes a shell command interactively with stdin/stdout/stderr connected.
// This is useful for opening editors or other interactive programs.
func RunShellCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
