package shell

import (
	"fmt"
	"log"
	"os/exec"
)

func IsProgramAvailable(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func MustExecuteProgram(name string, args ...string) string {
	// TODO: There's no error when executed with wrong software
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()

	if exitErr, ok := err.(*exec.ExitError); ok {
		log.Fatalln(fmt.Errorf("exit code %d: %s", exitErr.ExitCode(), output))
	}

	return string(output)
}
