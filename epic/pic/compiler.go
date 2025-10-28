package pic

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
	"epic/fs"
	"epic/utils"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func CompilePIC() {
	compileCore()
	fmt.Println()

	compileModules()
	fmt.Println()
}

func compileCore() {
	/*
		Compile "project/core/**" directory.
	*/
	cli.LogInfo("Compiling core...")

	coreDir := filepath.Join(ctx.CompilePIC.ProjectPath, "core")
	__compileDirectory(coreDir)

	cli.LogOk("Core compiled!")
}

func compileModules() {
	/*
		Compile "project/modules/<name>/**" directories.
	*/
	modules := fs.GetChildDirs(filepath.Join(ctx.CompilePIC.ProjectPath, "modules"))

	for _, module := range modules {
		if len(ctx.CompilePIC.Modules) > 0 && !slices.Contains(ctx.CompilePIC.Modules, module) {
			continue
		}

		cli.LogInfof("Compiling module: %s", module)

		moduleDir := filepath.Join(ctx.CompilePIC.ProjectPath, "modules", module)
		__compileDirectory(moduleDir)

		cli.LogOkf("Module compiled: %s", module)

	}

	// TODO: Add unknown module error
}

func __compileDirectory(inputPath string) {
	/*
		Compile all source files from :inputPath and put it into output/objects/**
		directories. The output structure of directory is mimicking :inputPath.
	*/

	absInputPath := fs.MustGetAbsPath(inputPath)

	for _, file := range fs.GetFilesByExtension(inputPath, ".c") {
		relPath, err := filepath.Rel(absInputPath, file.FullPath)
		if err != nil {
			cli.LogErrf("%v", err)
			os.Exit(1)
		}

		objectFileName := fs.ReplaceExtension(file.Name, ".o")
		// TODO: Fix this fucking path
		outputFile := filepath.Join(ctx.CompilePIC.OutputPath, "objects", filepath.Dir(relPath), objectFileName)

		fs.MustCreateDirTree(outputFile)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"--sysroot", inputPath,
			"-c", file.FullPath,
			"-o", outputFile,
			"-Os",
			"-std=c17",
			"-fPIC",
			"-nostdlib",
			"-nostartfiles",
			"-ffreestanding",
			"-fno-builtin",
			"-nodefaultlibs",
			"-ffunction-sections",
			"-fdata-sections",
			"-fno-ident",
			"-fno-jump-tables",
			"-falign-jumps=1",
			"-mgeneral-regs-only",
			"-fdiagnostics-color=always",
			"-fcf-protection=none",
			"-mno-sse",
			"-mno-mmx",
			"-mno-red-zone",
			"-mno-stack-arg-probe",
			"-fno-delete-null-pointer-checks",
			"-fno-asynchronous-unwind-tables",
			"-DPIC",
			"-ffixed-rbx",
		}

		// TODO: Add destination
		cli.LogInfof(" - Compiling: %s", filepath.Join(inputPath, relPath))

		output := utils.MingwGcc(params...)
		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
