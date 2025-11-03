package main

import (
	"epic/cmd"
	"os"
)

// TODO:
// - Implement custom symbols for compilation
// - Add author and project banner in project template
// - Add README in project template
// - README: Rewrite chapter about SAVE_GLOBAL() and GET_GLOBAL();

func main() {
	cmd.Run()

	os.Exit(0)
}
