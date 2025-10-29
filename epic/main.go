package main

import (
	"epic/cmd"
	"os"
)

// TODO:
// - Check standalone with printf() function
// - Add cool README
// - Init command - drop project structure to the disk
// - C++ Support:
// 		- Features which requires C++ runtime
//		- Can I use Mingw-GCC instead of G++? Would be easier.

// TODO:
// - Automatically obtain address of start and end of payload (in memory)
//

func main() {
	cmd.Execute()

	os.Exit(0)
}
