package main

import (
	"epic/cmd"
	"os"
)

// TODO:
// - Check standalone with printf() function
// - Add cool README
// - Init command - drop project structure to the disk

// TODO:
// - Automatically obtain address of start and end of payload (in memory)
//
// TODO Long-term: C++ Support?

func main() {
	cmd.Execute()

	os.Exit(0)
}
