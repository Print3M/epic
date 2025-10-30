package main

import (
	"epic/cmd"
	"os"
)

// TODO:
// - Add cool README
// - Check compatibility with Windows (default Mingw-w64 installation)

// TODO:
// - Automatically obtain address of start and end of payload (in memory)

func main() {
	cmd.Execute()

	os.Exit(0)
}
