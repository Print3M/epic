package main

import (
	"epic/cmd"
	"os"
)

// TODO:
// - Check standalone with printf() function
// - Add fancy banner (like in DllShimmer)
// - Add cool README
// - Check debug information

func main() {
	cmd.Execute()

	os.Exit(0)
}
