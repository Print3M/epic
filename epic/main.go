package main

import (
	"epic/builder"
	"epic/cli"
	"epic/ctx"
	"epic/shell"
	"fmt"
	"log"
)

/*
	TODO:
	1. Move from MinGW to GCC for PIC compilation (pure-GNU chain)
	2. Check dead code elimination capabilities
	3. Write it down why we switched from MinGW to GCC
	4. Fix calling convention
*/

func main() {
	cli.InitCLI()

	if !shell.IsProgramAvailable(ctx.CompilerPath) {
		log.Fatalf("Compiler not found: %s\n", ctx.CompilerPath)
	}

	if !shell.IsProgramAvailable(ctx.LinkerPath) {
		log.Fatalf("Linker not found: %s\n", ctx.LinkerPath)
	}

	if !ctx.NoPIC {
		// Compile PIC payload
		fmt.Println()
		builder.BuildPIC()
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
*/
