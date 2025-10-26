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
	5. Leave MinGW for standalone and loader compilation

	TODO: Try for x86_64-w64-mingw32-ld
*/

func main() {
	cli.InitCLI()

	if !ctx.NoPIC {
		if !shell.IsProgramAvailable(ctx.GccPath) {
			log.Fatalf("GCC compiler not found: %s\n", ctx.GccPath)
		}

		if !shell.IsProgramAvailable(ctx.LinkerPath) {
			log.Fatalf("Linker not found: %s\n", ctx.LinkerPath)
		}

		// Compile PIC payload
		fmt.Println()
		builder.BuildPIC()
	}

	if !ctx.NoLoader {
		if !shell.IsProgramAvailable(ctx.MingwGccPath) {
			log.Fatalf("MinGW-GCC compiler not found: %s\n", ctx.MingwGccPath)
		}

		// Compile loader
		fmt.Println()
		builder.BuildLoader()
	}

	if !ctx.NoStandalone {
		if !shell.IsProgramAvailable(ctx.MingwGccPath) {
			log.Fatalf("MinGW-GCC compiler not found: %s\n", ctx.MingwGccPath)
		}

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
