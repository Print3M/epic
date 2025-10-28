package main

import (
	"epic/cmd"
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
// - Maybe there should be two modes: build and link? I need some option to just link object files.

func main() {
	cmd.Execute()

	os.Exit(0)
}
