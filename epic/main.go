package main

import (
	"epic/builder"
	"epic/cli"
	"epic/ctx"
	"epic/shell"
	"fmt"
	"os"
)

// TODO:
// Before run...
// - Clean output/objects/
// - Clean output/assets/
// TODO:
// - Check standalone with printf() function
// - Add fancy banner (like in DllShimmer)
// - Print nice output with generated files and what to do next with them
// - Test standalone with printf()
// - Add cool README

func assertCompiler() {
	if !shell.IsProgramAvailable(ctx.CompilerPath) {
		cli.LogErr(fmt.Sprintf("Mingw-w64 GCC compiler not found: %s\n", ctx.CompilerPath))
		os.Exit(1)
	}
}

func assertLinker() {
	if !shell.IsProgramAvailable(ctx.LinkerPath) {
		cli.LogErr(fmt.Sprintf("Mingw-w64 ld linker not found: %s\n", ctx.LinkerPath))
		os.Exit(1)
	}
}

func assertObjcopy() {
	if !shell.IsProgramAvailable(ctx.LinkerPath) {
		cli.LogErr(fmt.Sprintf("Mingw-w64 objcopy tool not found: %s\n", ctx.ObjcopyPath))
		os.Exit(1)
	}
}

func main() {
	cli.InitCLI()

	if ctx.Debug {
		cli.LogDbg("Debug mode enabled")
	}

	if !ctx.NoPIC {
		assertCompiler()
		assertLinker()
		assertObjcopy()

		// Compile PIC payload
		builder.BuildPIC()
	}

	if !ctx.NoLoader {
		assertCompiler()

		// Compile loader
		fmt.Println()
		builder.BuildLoader()
	}

	if !ctx.NoNonPIC {
		assertCompiler()

		// Compile standalone
		fmt.Println()
		builder.BuildNonPIC()
	}

	os.Exit(0)
}
