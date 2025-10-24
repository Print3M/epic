package shell

import (
	"fmt"
	"os/exec"
)

func IsProgramInstalled(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func ExecuteProgram(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	if exitErr, ok := err.(*exec.ExitError); ok {
		return string(output), fmt.Errorf("exit code %d: %s", exitErr.ExitCode(), output)
	}

	return string(output), err
}
