package main

import (
	"epic/builder"
	"epic/cli"
	"fmt"
)

func main() {
	flags := cli.ParseCli()

	/*
		TODO: There's no error when executed with wrong software
		if !shell.IsProgramAvailable(flags.CompilerPath) {
			log.Fatalf("Compiler not found: %s\n", flags.CompilerPath)
		}

		if !shell.IsProgramAvailable(flags.LinkerPath) {
			log.Fatalf("Linker not found: %s\n", flags.LinkerPath)
		}
	*/

	if !flags.NoPIC {
		// Compile PIC payload
		fmt.Println()
		builder.BuildPIC(flags, []string{"pwd"})
	}

	if !flags.NoLoader {
		// Compile loader
		fmt.Println()
		builder.BuildLoader(flags)
	}

	if !flags.NoStandalone {
		// Compile standalone
		fmt.Println()
		builder.BuildStandalone(flags)
	}
}

/*
	TODO:
		- Add colors and loading bars in CLI
		- Add global context
		- Refactor
*/
