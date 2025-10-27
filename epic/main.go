package main

import (
	"epic/builder"
	"epic/cli"
	"epic/ctx"
	"epic/shell"
	"fmt"
	"log"
	"os"
)

// TODO:
// Before run...
// - Clean output/objects/
// - Clean output/assets/
// TODO:
// - Check standalone with printf() function
// - Add colors and loading bars in CLI
// - Add fancy banner (like in DllShimmer)
// - Print nice output with generated files and what to do next with them

func main() {
	cli.InitCLI()

	if ctx.Debug {
		fmt.Println("[DBG] Debug mode enabled")
	}

	if !ctx.NoPIC {
		if !shell.IsProgramAvailable(ctx.CompilerPath) {
			log.Fatalf("GCC compiler not found: %s\n", ctx.CompilerPath)
		}

		if !shell.IsProgramAvailable(ctx.LinkerPath) {
			log.Fatalf("Linker not found: %s\n", ctx.LinkerPath)
		}

		// Compile PIC payload
		fmt.Println()
		builder.BuildPIC()
	}

	if !ctx.NoLoader {
		if !shell.IsProgramAvailable(ctx.CompilerPath) {
			log.Fatalf("MinGW-GCC compiler not found: %s\n", ctx.CompilerPath)
		}

		// Compile loader
		fmt.Println()
		builder.BuildLoader()
	}

	if !ctx.NoStandalone {
		if !shell.IsProgramAvailable(ctx.CompilerPath) {
			log.Fatalf("MinGW-GCC compiler not found: %s\n", ctx.CompilerPath)
		}

		// Compile standalone
		fmt.Println()
		builder.BuildStandalone()
	}

	os.Exit(0)
}
