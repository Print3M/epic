package main

import (
	"epic/builder"
	"epic/cli"
	"epic/ctx"
	"fmt"
)

func main() {
	cli.InitCLI()

	/*
		TODO: There's no error when executed with wrong software
		if !shell.IsProgramAvailable(flags.CompilerPath) {
			log.Fatalf("Compiler not found: %s\n", flags.CompilerPath)
		}

		if !shell.IsProgramAvailable(flags.LinkerPath) {
			log.Fatalf("Linker not found: %s\n", flags.LinkerPath)
		}
	*/

	if !ctx.NoPIC {
		// Compile PIC payload
		fmt.Println()
		builder.BuildPIC([]string{"pwd"})
	}

	if !ctx.NoLoader {
		// Compile loader
		fmt.Println()
		builder.BuildLoader()
	}

	if !ctx.NoStandalone {
		// Compile standalone
		fmt.Println()
		builder.BuildStandalone()
	}
}

/*
	TODO:
		- Add colors and loading bars in CLI
		- Add fancy banner (like in DllShimmer)
		- Add global context & Cobra CLI
		- Add dynamic modules
		- Refactor
*/
