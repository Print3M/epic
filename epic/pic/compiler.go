package pic

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
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

	__compileProjectDirectory("core")

	cli.LogOk("Core compiled!")
}

func compileModules() {
	/*
		Compile "project/modules/<name>/**" directories.
	*/
	modules := utils.GetChildDirs(filepath.Join(ctx.CompilePIC.ProjectPath, "modules"))

	for _, module := range modules {
		if !slices.Contains(ctx.CompilePIC.Modules, module) {
			continue
		}

		cli.LogInfof("Compiling '%s' module...", module)

		moduleDir := filepath.Join("modules", module)
		__compileProjectDirectory(moduleDir)

		cli.LogOkf("Module '%s' compiled!", module)

	}

	// TODO: Add unknown module error
}

func __compileProjectDirectory(projectDir string) {
	/*
		Compile all source files from :inputPath and put it into output/objects/**
		directories. The output structure of directory is mimicking :inputPath.
	*/

	absProjectDir := utils.MustGetAbsPath(filepath.Join(ctx.CompilePIC.ProjectPath, projectDir))

	for _, file := range utils.GetFilesByExtension(absProjectDir, ".c") {
		relPath, err := filepath.Rel(absProjectDir, file.FullPath)
		if err != nil {
			cli.LogErrf("%v", err)
			os.Exit(1)
		}

		inputFile := filepath.Join(ctx.CompilePIC.ProjectPath, projectDir, file.Name)
		outputDir := filepath.Join(
			ctx.CompilePIC.OutputPath,
			projectDir,
			filepath.Dir(relPath),
		)
		utils.MustCreateDirTree(outputDir)
		outputFile := filepath.Join(outputDir, utils.ReplaceExtension(file.Name, ".o"))

		cli.LogInfof(" ‣ %s -> %s", inputFile, outputFile)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"--sysroot", projectDir,
			"-c", file.FullPath,
			"-o", outputFile,
			"-Os",
			"-std=c17",
			"-fPIC",
			"-nostdlib",
			"-nostdinc",
			"-nostartfiles",
			"-ffreestanding",
			"-fno-builtin",
			"-nodefaultlibs",
			"-ffunction-sections",
			"-fdata-sections",
			"-fno-ident",
			"-fno-rtti",
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
			"-I", filepath.Join(ctx.CompilePIC.ProjectPath, "include"),
			"-I", ctx.CompilePIC.ProjectPath,
		}

		output := utils.MingwGcc(params...)
		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
