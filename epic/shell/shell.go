package shell

import (
	"epic/cli"
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
		cli.LogDbg("Command executed:")
		fmt.Println(cmd.String())
	}

	output, err := cmd.CombinedOutput()

	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		cli.LogErr(fmt.Sprintf("'%s' exit code: %d", name, exitCode))
		log.Fatalln(string(output))
	}

	return string(output)
}
