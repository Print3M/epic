package builder

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
	"epic/fs"
	"epic/shell"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed assets/linker.ld
var linkerScriptContent string

func BuildPIC() {
	buildCore()
	fmt.Println()

	buildModules()
	fmt.Println()

	linkedExecutable := linkExecutable()
	extractTextSection(linkedExecutable)
}

func extractTextSection(file string) {
	/*
		Using objcopy tool extract .text section from compiled executable.
	*/
	cli.LogInfo("Extracting '.text' section from PIC-PE...")
	outputFile := fs.OutputPath("payload.bin")

	output := shell.MustExecuteProgram(ctx.ObjcopyPath, "-O", "binary", "--only-section=.text", file, outputFile)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("PIC payload extracted!")
}

func linkExecutable() string {
	/*
		Using ld linker link all object files
	*/

	var modules []string
	for _, m := range fs.GetOutputModules() {
		modules = append(modules, m.Name)
	}
	cli.LogInfo(fmt.Sprintf("Linking PIC core + modules (%s)", strings.Join(modules, ",")))

	outputPeFile := fs.OutputPath("assets", "payload.exe")
	linkerMapFile := fs.OutputPath("assets", "payload.linker.map")
	fs.MustCreateDirTree(outputPeFile)

	linkerScriptFile := fs.OutputPath("assets", "linker.ld")
	fs.MustCreateDirTree(linkerScriptFile)

	fs.MustWriteFile(linkerScriptFile, linkerScriptContent)

	objectFiles := getObjectFiles()

	if ctx.Debug {
		cli.LogDbg("Linking:")
		for _, f := range objectFiles {
			cli.LogDbg(fmt.Sprintf(" - Linking: %s", f))
		}
	}

	params := []string{
		"--gc-sections",
		"-Map", linkerMapFile,
		"-T", linkerScriptFile,
		"-o", outputPeFile,
		"--image-base=0x00",
	}

	params = append(params, objectFiles...)

	output := shell.MustExecuteProgram(ctx.LinkerPath, params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("PIC-PE linked!")

	return outputPeFile
}

func getObjectFiles() []string {
	var objectFiles []string

	// Collecting core
	for _, f := range fs.GetFilesByExtension(fs.OutputPath("objects", "core"), ".o") {
		objectFiles = append(objectFiles, f.FullPath)
	}

	// Collecting modules
	for _, m := range fs.GetOutputModules() {
		for _, f := range fs.GetFilesByExtension(m.Path, ".o") {
			objectFiles = append(objectFiles, f.FullPath)
		}
	}

	return objectFiles
}

func buildCore() {
	/*
		Build "project/core/**" directory.
	*/
	cli.LogInfo("Building core...")

	buildDirectory(fs.ProjectPath("core"), 1)

	cli.LogOk("Core built!")
}

func buildModules() {
	/*
		Build "project/modules/<name>/**" directories.
	*/
	modules := fs.GetProjectModules()

	for _, module := range modules {
		cli.LogInfo(fmt.Sprintf("Building module: %s", module.Name))

		buildDirectory(fs.ProjectPath("modules", module.Name), 2)

		cli.LogOk(fmt.Sprintf("Module built: %s", module.Name))

	}
}

func buildDirectory(rootDir string, logIndent int) {
	/*
		Build all source files from :rootDir and put it into output/objects/**
		directories. The output structure of directory is mimicking the rootDir.
	*/
	for _, file := range fs.GetFilesByExtension(rootDir, ".c") {
		relPath, err := filepath.Rel(ctx.ProjectPath, file.FullPath)
		if err != nil {
			cli.LogErr(fmt.Sprintf("%v", err))
			os.Exit(1)
		}

		objectFileName := fs.ReplaceExtension(file.Name, ".o")
		outputPath := fs.OutputPath("objects", filepath.Dir(relPath), objectFileName)

		fs.MustCreateDirTree(outputPath)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"--sysroot", rootDir,
			"-c", file.FullPath,
			"-o", outputPath,
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

		cli.LogInfo(fmt.Sprintf(" - Building: %s", file.FullPath))

		output := shell.MustExecuteProgram("x86_64-w64-mingw32-gcc", params...)

		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
