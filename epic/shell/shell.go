package shell

import (
	"epic/ctx"
	"fmt"
	"log"
	"os/exec"
)

func IsProgramAvailable(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func MustExecuteProgram(name string, args ...string) string {
	cmd := exec.Command(name, args...)

	if ctx.Debug {
		fmt.Println("[DBG] Command executed:", cmd.String())
	}

	output, err := cmd.CombinedOutput()

	if exitErr, ok := err.(*exec.ExitError); ok {
		log.Fatalln(fmt.Errorf("exit code %d: %s", exitErr.ExitCode(), output))
	}

	return string(output)
}
