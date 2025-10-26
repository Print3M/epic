package cli

import (
	"epic/ctx"
	"epic/fs"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func MustValidateProjectPath(projectPath string) {
	if p := filepath.Join(projectPath, "core"); !fs.EntryExists(p) {
		log.Fatalf("Invalid project structure. Path doesn't exist: %s", p)
	}

	if p := filepath.Join(projectPath, "core", "main.c"); !fs.EntryExists(p) {
		log.Fatalf("Invalid project structure. Path doesn't exist: %s", p)
	}

	if p := filepath.Join(projectPath, "modules"); !fs.EntryExists(p) {
		log.Fatalf("Invalid project structure. Path doesn't exist: %s", p)
	}
}

func InitCLI() {
	flag.StringVar(&ctx.ProjectPath, "p", "", "")
	flag.StringVar(&ctx.ProjectPath, "project", "", "")

	flag.StringVar(&ctx.OutputPath, "o", "", "")
	flag.StringVar(&ctx.OutputPath, "output", "", "")

	flag.StringVar(&ctx.GccPath, "gcc", "", "")
	flag.StringVar(&ctx.LinkerPath, "ld", "", "")
	flag.StringVar(&ctx.MingwGccPath, "mingw-gcc", "", "")

	flag.BoolVar(&ctx.NoPIC, "no-pic", false, "")
	flag.BoolVar(&ctx.NoLoader, "no-loader", false, "")
	flag.BoolVar(&ctx.NoStandalone, "no-standalone", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: epic -p <path> -o <path>\n")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Printf("  %-26s %s\n", "-p, --project <path>", "Project directory (required)")
		fmt.Printf("  %-26s %s\n", "-o, --output <path>", "Output directory (required)")
		fmt.Printf("  %-26s %s\n", "--no-pic", "Disable PIC payload building")
		fmt.Printf("  %-26s %s\n", "--no-loader", "Disable loader building")
		fmt.Printf("  %-26s %s\n", "--no-standalone", "Disable standalone building")
		fmt.Println()
		fmt.Println(" Advanced parameters:")
		fmt.Printf("  %-26s %s\n", "--ld <path>", "Path to LD (GNU) linker")
		fmt.Printf("  %-26s %s\n", "--gcc <path>", "Path to GCC compiler")
		fmt.Printf("  %-26s %s\n", "--mingw-gcc <path>", "Path to MinGW-GCC compiler. It's not used for PIC payload compilation.")
		fmt.Printf("  %-26s %s\n", "--mingw-gcc <path>", "Path to MinGW-GCC compiler. It's not used for PIC payload compilation.")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  epic -p project/ -o output/")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if ctx.ProjectPath == "" || ctx.OutputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx.ProjectPath = fs.MustGetAbsPath(ctx.ProjectPath)
	ctx.OutputPath = fs.MustGetAbsPath(ctx.OutputPath)

	if ctx.MingwGccPath == "" {
		ctx.MingwGccPath = "x86_64-w64-mingw32-gcc"
	}

	if ctx.GccPath == "" {
		ctx.GccPath = "gcc"
	}

	if ctx.LinkerPath == "" {
		ctx.LinkerPath = "ld"
	}

	if ctx.NoPIC && ctx.NoLoader && ctx.NoStandalone {
		fmt.Println("You've disabled everything. I can't offer you anything more...")
		os.Exit(1)
	}

	MustValidateProjectPath(ctx.ProjectPath)

	// TODO: Implement it as parameter
	ctx.Modules = []string{"pwd"}
}
