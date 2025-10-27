package builder

import (
	_ "embed"
	"epic/ctx"
	"epic/fs"
	"epic/shell"
	"fmt"
	"log"
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
	outputFile := fs.OutputPath("payload.bin")

	output := shell.MustExecuteProgram(ctx.ObjcopyPath, "-O", "binary", "--only-section=.text", file, outputFile)
	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[+] PIC payload extracted from ELF.")
}

func linkExecutable() string {
	/*
		Using ld linker link all object files
	*/

	// TODO: Print which modules are being linked...
	fmt.Println("[*] Linking PIC payload...")

	outputElfFile := fs.OutputPath("payload.elf")
	linkerScriptFile := fs.OutputPath("assets", "linker.ld")

	fs.MustCreateDirTree(linkerScriptFile)
	fs.MustWriteFile(linkerScriptFile, linkerScriptContent)

	linkerMapFile := fs.OutputPath("payload.linker.map")

	objectFiles := getObjectFiles()

	params := []string{
		// "--print-gc-sections", // TODO: Debug option
		"--gc-sections",
		"-Map", linkerMapFile,
		"-T", linkerScriptFile,
		"-o", outputElfFile,
	}

	params = append(params, objectFiles...)

	output := shell.MustExecuteProgram(ctx.LinkerPath, params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[+] PIC ELF linked.")

	return outputElfFile
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
	fmt.Println("[*] Building core...")

	buildDirectory(fs.ProjectPath("core"), 1)

	fmt.Println("[+] Core built!")
}

func buildModules() {
	/*
		Build "project/modules/<name>/**" directories.
	*/
	fmt.Println("[*] Building modules...")

	modules := fs.GetProjectModules()

	for _, module := range modules {
		fmt.Println("\t[*] Building module:", module.Name)

		buildDirectory(fs.ProjectPath("modules", module.Name), 2)

		fmt.Println("\t[+] Module built:", module.Name)
	}

	fmt.Println("[+] Modules built!")
}

func buildDirectory(rootDir string, logIndent int) {
	/*
		Build all source files from :rootDir and put it into output/objects/**
		directories. The output structure of directory is mimicking the rootDir.
	*/
	for _, file := range fs.GetFilesByExtension(rootDir, ".c") {
		relPath, err := filepath.Rel(ctx.ProjectPath, file.FullPath)
		if err != nil {
			log.Println(err)
		}

		objectFileName := fs.ReplaceExtension(file.Name, ".o")
		outputPath := fs.OutputPath("objects", filepath.Dir(relPath), objectFileName)

		fs.MustCreateDirTree(outputPath)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"--sysroot", rootDir,
			"-c", file.FullPath,
			"-o", outputPath,
			"-nostdlib",
			"-fPIC",
			"-nostartfiles",
			"-Os",
			"-fno-asynchronous-unwind-tables",
			"-ffreestanding",
			"-fno-builtin",
			"-ffunction-sections",
			"-fdata-sections",
			"-fno-ident",
			"-falign-jumps=1",
			"-mno-sse",
			"-mno-mmx",
			"-mgeneral-regs-only",
			"-mno-stack-arg-probe",
			"-mno-red-zone",
			"-fdiagnostics-color=always",
			"-fcf-protection=none",
			"-fno-delete-null-pointer-checks",
			"-std=c17",
			"-nodefaultlibs",
		}

		fmt.Println(strings.Repeat("\t", logIndent), file.FullPath)

		output := shell.MustExecuteProgram("x86_64-w64-mingw32-gcc", params...)

		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
