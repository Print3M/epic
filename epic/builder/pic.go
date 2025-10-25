package builder

import (
	_ "embed"
	"epic/cli"
	"epic/fs"
	"epic/shell"
	"fmt"
	"path/filepath"
	"strings"
)

//go:embed assets/linker.ld
var linkerScriptContent string

func BuildPIC(flags *cli.CliFlags, modules []string) {
	buildCore(flags)
	fmt.Println()

	buildModules(flags)
	fmt.Println()

	fmt.Println("[*] Linking PIC payload...")

	var objectFiles []string

	// Collect core object files
	corePath := filepath.Join(flags.OutputPath, "objects/core")
	for _, e := range fs.GetFilesByExtension(corePath, ".o") {
		objectFiles = append(objectFiles, e.FullPath)
	}

	// Collect modules object files
	for _, module := range modules {
		modulePath := filepath.Join(flags.OutputPath, "objects/modules", module)

		for _, e := range fs.GetFilesByExtension(modulePath, ".o") {
			objectFiles = append(objectFiles, e.FullPath)
		}
	}

	outputFile := filepath.Join(flags.OutputPath, "payload.bin")

	linkerScriptFile := filepath.Join(flags.OutputPath, "assets", "linker.ld")
	fs.CreateDirTree(linkerScriptFile)
	fs.MustWriteFile(linkerScriptFile, linkerScriptContent)

	params := []string{
		"-T", linkerScriptFile,
		"-o", outputFile,
	}

	params = append(params, objectFiles...)

	output := shell.MustExecuteProgram(flags.LinkerPath, params...)

	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[+] PIC payload linked!")
}

func buildCore(flags *cli.CliFlags) {
	fmt.Println("[*] Building core...")

	buildDirectory(flags, "core", 1)

	fmt.Println("[+] Core built!")
}

type Module struct {
	Name     string
	FullPath string
}

func getModuleNames(modulesPath string) []Module {
	// Each module is a separate directory in 'modules/' path
	entries := fs.GetDirectories(modulesPath)

	var modules []Module

	for _, entry := range entries {
		if !entry.IsDir {
			continue
		}

		modules = append(modules, Module{
			Name:     entry.Name,
			FullPath: entry.FullPath,
		})
	}

	return modules
}

func buildModules(flags *cli.CliFlags) {
	fmt.Println("[*] Building modules...")

	modules := getModuleNames(filepath.Join(flags.InputPath, "modules"))

	for _, module := range modules {
		fmt.Println("\t[*] Building module:", module.Name)

		buildDirectory(flags, filepath.Join("modules/", module.Name), 2)

		fmt.Println("\t[+] Module built:", module.Name)
	}

	fmt.Println("[+] Modules built!")
}

func buildDirectory(flags *cli.CliFlags, dir string, logIndent int) {
	corePath := filepath.Join(flags.InputPath, dir)

	for _, source := range fs.GetFilesByExtension(corePath, ".c") {
		outputRelPath := fs.ReplaceExtension(source.RelPath, ".o")
		outputFullPath := filepath.Join(flags.OutputPath, "objects", dir, outputRelPath)

		fs.CreateDirTree(outputFullPath)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"--sysroot",
			corePath,
			"-c",
			source.FullPath,
			"-o",
			outputFullPath,
			"-nostdlib",
			"-fPIC",
			"-nostartfiles",
			"-Os",
			"-fno-asynchronous-unwind-tables",
			"-ffreestanding",
			"-fno-builtin",
			"-ffunction-sections",
			"-fno-ident",
			"-falign-jumps=1",
			"-mno-sse",
			"-mno-mmx",
			"-mgeneral-regs-only",
			"-mno-stack-arg-probe",
			"-mno-red-zone",
			"-fdiagnostics-color=always",
			"-std=c17",
		}

		fmt.Println(strings.Repeat("\t", logIndent), source.FullPath)

		output := shell.MustExecuteProgram(flags.CompilerPath, params...)

		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
