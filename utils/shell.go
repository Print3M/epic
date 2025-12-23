package utils

import (
	"epic/cli"
	"epic/ctx"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func IsProgramAvailable(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
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

// TODO: Implement Windows compatibility

func MingwLd(params ...string) string {
	ld := ctx.MingwLdPath
	if ld == "" {
		ld = "x86_64-w64-mingw32-ld"

		if IsWindows() {
			ld = "ld.exe"
		}
	}

	if !IsProgramAvailable(ld) {
		cli.LogErrf("Mingw-w64 ld linker not found: %s", ld)
		cli.LogErrf("Specify path to Mingw-w64 ld: `--mingw-w64-ld <path>`")
		os.Exit(1)
	}

	output := MustExecuteProgram(ld, params...)

	return output
}

func MingwGcc(params ...string) string {
	gcc := ctx.MingwGccPath
	if gcc == "" {
		gcc = "x86_64-w64-mingw32-gcc"

		if IsWindows() {
			gcc = "gcc.exe"
		}
	}

	if !IsProgramAvailable(gcc) {
		cli.LogErrf("Mingw-w64 GCC compiler not found: %s", gcc)
		cli.LogErrf("Specify path to Mingw-w64 GCC: `--mingw-w64-gcc <path>`")
		os.Exit(1)
	}

	output := MustExecuteProgram(gcc, params...)

	return output
}

func MingwObjcopy(params ...string) string {
	objcopy := ctx.MingwObjcopyPath
	if objcopy == "" {
		objcopy = "x86_64-w64-mingw32-objcopy"

		if IsWindows() {
			objcopy = "objcopy.exe"
		}
	}

	if !IsProgramAvailable(objcopy) {
		cli.LogErrf("Mingw-w64 objcopy tool not found: %s", objcopy)
		cli.LogErrf("Specify path to Mingw-w64 objcopy: `--mingw-w64-objcopy <path>`")
		os.Exit(1)
	}

	output := MustExecuteProgram(objcopy, params...)

	return output
}
