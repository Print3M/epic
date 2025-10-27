package cli

import (
	"epic/ctx"
	"epic/fs"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func mustValidateProjectPath(projectPath string) {
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

func parseModules(param string) {
	if len(param) == 0 {
		return
	}

	// Parse comma separated list of modules
	modules := strings.Split(param, ",")

	for i := range modules {
		modules[i] = strings.ToLower(strings.TrimSpace(modules[i]))
	}

	ctx.Modules = modules
}

func InitCLI() {
	flag.StringVar(&ctx.ProjectPath, "p", "", "")
	flag.StringVar(&ctx.ProjectPath, "project", "", "")

	flag.StringVar(&ctx.OutputPath, "o", "", "")
	flag.StringVar(&ctx.OutputPath, "output", "", "")

	var modules string
	flag.StringVar(&modules, "m", "", "")
	flag.StringVar(&modules, "modules", "", "")

	flag.BoolVar(&ctx.Debug, "debug", false, "")

	flag.StringVar(&ctx.CompilerPath, "mingw-w64-gcc", "", "")
	flag.StringVar(&ctx.LinkerPath, "mingw-w64-ld", "", "")
	flag.StringVar(&ctx.ObjcopyPath, "mingw-w64-objcopy", "", "")

	flag.BoolVar(&ctx.NoPIC, "no-pic", false, "")
	flag.BoolVar(&ctx.NoLoader, "no-loader", false, "")
	flag.BoolVar(&ctx.NoStandalone, "no-standalone", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: epic -p <path> -o <path> -m pwd,ls\n")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println()
		fmt.Printf("  %-27s %s\n", "-p, --project <path>", "Project directory (required)")
		fmt.Printf("  %-27s %s\n", "-o, --output <path>", "Output directory (required)")
		fmt.Printf("  %-27s %s\n", "-m, --modules <path>", "Included module names (comma-separated list, default: none)")
		fmt.Printf("  %-27s %s\n", "--no-pic", "Disable PIC payload building")
		fmt.Printf("  %-27s %s\n", "--no-loader", "Disable loader building")
		fmt.Printf("  %-27s %s\n", "--no-standalone", "Disable standalone building")
		fmt.Printf("  %-27s %s\n", "--debug", "Enable verbose debugging")
		fmt.Println()
		fmt.Println("Toolchain options:")
		fmt.Printf("  %-27s %s\n", "--mingw-w64-gcc <path>", "Path to Mingw-w64 GCC compiler")
		fmt.Printf("  %-27s %s\n", "--mingw-w64-ld <path>", "Path to Mingw-w64 LD linker")
		fmt.Printf("  %-27s %s\n", "--mingw-w64-objcopy <path>", "Path to Mingw-w64 objcopy tool")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println()
		fmt.Println("  epic -p project/ -o output/")
		fmt.Println()
		fmt.Println("Created by Print3M (print3m.github.io)")
		fmt.Println()
	}

	flag.Parse()

	if ctx.NoPIC && ctx.NoLoader && ctx.NoStandalone {
		fmt.Println("You've disabled everything. I can't offer you anything more...")
		os.Exit(1)
	}

	if ctx.ProjectPath == "" || ctx.OutputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx.ProjectPath = fs.MustGetAbsPath(ctx.ProjectPath)
	ctx.OutputPath = fs.MustGetAbsPath(ctx.OutputPath)

	if ctx.CompilerPath == "" {
		ctx.CompilerPath = "x86_64-w64-mingw32-gcc"
	}

	if ctx.LinkerPath == "" {
		// ctx.LinkerPath = "ld"
		// TODO: Check if this works as previous `ld`
		ctx.LinkerPath = "x86_64-w64-mingw32-ld"
	}

	if ctx.ObjcopyPath == "" {
		ctx.ObjcopyPath = "x86_64-w64-mingw32-objcopy"
	}

	parseModules(modules)

	mustValidateProjectPath(ctx.ProjectPath)
}
