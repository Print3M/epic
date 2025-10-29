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
// - Implement minimalistic include/win32 and include/libc
// - Get rid of any implicit includes from the source code
// - Change root of include paths in GCC
// - Implement string.h functions
// - Check if these headers are always included (even with -nostdinc):
//   <float.h>, <iso646.h>, <limits.h>, <stdarg.h>, <stdbool.h>, <stddef.h>, <stdint.h>

func main() {
	cmd.Execute()

	os.Exit(0)
}
