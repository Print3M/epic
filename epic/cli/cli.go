package cli

import (
	"epic/fs"
	"flag"
	"fmt"
	"os"
)

type CliFlags struct {
	InputPath    string
	OutputPath   string
	CompilerPath string
}

func ParseCli() *CliFlags {
	var flags CliFlags

	flag.StringVar(&flags.InputPath, "i", "", "")
	flag.StringVar(&flags.InputPath, "input", "", "")

	flag.StringVar(&flags.OutputPath, "o", "", "")
	flag.StringVar(&flags.OutputPath, "output", "", "")

	flag.StringVar(&flags.CompilerPath, "gcc", "", "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: Epic -i <path> -o <path> -p <path>\n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-26s %s\n", "-i, --input <path>", "Input directory (required)")
		fmt.Printf("  %-26s %s\n", "-o, --output <path>", "Output directory (required)")
		fmt.Printf("  %-26s %s\n", "--gcc <path>", "Path to GCC (MinGW) compiler")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  DllShimmer -i version.dll -o ./project -x 'C:\\Windows\\System32\\version.dll' -m")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if flags.InputPath == "" || flags.OutputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	flags.InputPath = fs.MustGetAbsPath(flags.InputPath)
	flags.OutputPath = fs.MustGetAbsPath(flags.OutputPath)

	if flags.CompilerPath == "" {
		flags.CompilerPath = "x86_64-w64-mingw32-gcc"
	}

	// TODO: Check if input directory has a correct structure

	return &flags
}
