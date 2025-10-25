package builder

import (
	_ "embed"
	"epic/ctx"
	"epic/fs"
	"epic/shell"
	"fmt"
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

	fmt.Println("[*] Linking PIC payload...")

	var objectFiles []string

	// Collect core object files
	corePath := fs.OutputPath("objects", "core")
	for _, e := range fs.GetFilesByExtension(corePath, ".o") {
		objectFiles = append(objectFiles, e.FullPath)
	}

	// Collect modules object files
	for _, module := range ctx.Modules {
		modulePath := fs.OutputPath("objects", "modules", module)

		for _, e := range fs.GetFilesByExtension(modulePath, ".o") {
			objectFiles = append(objectFiles, e.FullPath)
		}
	}

	// TODO: Return to payload.bin
	outputFile := fs.OutputPath("payload.elf")

	linkerScriptFile := fs.OutputPath("assets", "linker.ld")
	fs.CreateDirTree(linkerScriptFile)
	fs.MustWriteFile(linkerScriptFile, linkerScriptContent)

	linkerMapFile := fs.OutputPath("payload.linker.map")

	params := []string{
		"--print-gc-sections",
		"--gc-sections",
		"--entry=main",
		"-Map", linkerMapFile,
		"-T", linkerScriptFile,
		"-o", outputFile,
	}

	params = append(params, objectFiles...)

	output := shell.MustExecuteProgram(ctx.LinkerPath, params...)

	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[+] PIC payload linked!")
}

func buildCore() {
	fmt.Println("[*] Building core...")

	buildDirectory("core", 1)

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

func buildModules() {
	fmt.Println("[*] Building modules...")

	modules := getModuleNames(fs.ProjectPath("modules"))

	for _, module := range modules {
		fmt.Println("\t[*] Building module:", module.Name)

		buildDirectory(filepath.Join("modules", module.Name), 2)

		fmt.Println("\t[+] Module built:", module.Name)
	}

	fmt.Println("[+] Modules built!")
}

func buildDirectory(dir string, logIndent int) {
	corePath := fs.ProjectPath(dir)

	for _, source := range fs.GetFilesByExtension(corePath, ".c") {
		outputRelPath := fs.ReplaceExtension(source.RelPath, ".o")
		outputFullPath := fs.OutputPath("objects", dir, outputRelPath)

		fs.CreateDirTree(outputFullPath)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"--sysroot", corePath,
			"-c", source.FullPath,
			"-o", outputFullPath,
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
			"-fdata-sections",
		}

		fmt.Println(strings.Repeat("\t", logIndent), source.FullPath)

		output := shell.MustExecuteProgram("gcc", params...)

		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
