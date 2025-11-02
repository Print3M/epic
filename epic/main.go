package main

import (
	"epic/cmd"
	"os"
)

// TODO:
// - Automatically obtain address of start and end of payload (in memory)
// - Implement custom symbols for compilation

func main() {
	cmd.Run()

	os.Exit(0)
}
