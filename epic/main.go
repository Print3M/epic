package main

import (
	"epic/build"
	"epic/cli"
	"epic/shell"
	"log"
)

func main() {
	flags := cli.ParseCli()

	if !shell.IsProgramInstalled(flags.CompilerPath) {
		log.Fatalf("Compiler '%s' not found", flags.CompilerPath)
	}

	build.BuildCore(flags)

	// TODO: Solid function for shell command execution
	// TODO: Check if mingw is installed, path to mingw
	// TODO: Compile core
	// TODO: Compile modules
}
