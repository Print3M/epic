package utils

import (
	"epic/cli"
	"epic/ctx"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func IsProgramAvailable(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func MustExecuteProgram(name string, args ...string) string {
	var validArgs []string
	for _, arg := range args {
		if len(strings.TrimSpace(arg)) > 0 {
			validArgs = append(validArgs, arg)
		}
	}

	cmd := exec.Command(name, validArgs...)

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

func MingwLd(params ...string) string {
	ld := ctx.MingwLdPath
	if ld == "" {
		ld = "x86_64-w64-mingw32-ld"
	}

	if !IsProgramAvailable(ld) {
		cli.LogErrf("Mingw-w64 ld linker not found: %s\n", ld)
		os.Exit(1)
	}

	output := MustExecuteProgram(ld, params...)

	return output
}

func MingwGcc(params ...string) string {
	gcc := ctx.MingwGccPath
	if gcc == "" {
		gcc = "x86_64-w64-mingw32-gcc"
	}

	if !IsProgramAvailable(gcc) {
		cli.LogErrf("Mingw-w64 GCC compiler not found: %s", gcc)
		os.Exit(1)
	}

	output := MustExecuteProgram(gcc, params...)

	return output
}

func MingwObjcopy(params ...string) string {
	objcopy := ctx.MingwObjcopyPath
	if objcopy == "" {
		objcopy = "x86_64-w64-mingw32-objcopy"
	}

	if !IsProgramAvailable(objcopy) {
		cli.LogErrf("Mingw-w64 objcopy tool not found: %s\n", objcopy)
		os.Exit(1)
	}

	output := MustExecuteProgram(objcopy, params...)

	return output
}
