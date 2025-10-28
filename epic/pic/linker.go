package pic

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
	"epic/fs"
	"epic/utils"
	"fmt"
	"path/filepath"
	"slices"
)

//go:embed assets/linker.ld
var linkerScriptContent string

func LinkPIC() {
	linkedExecutable := linkExecutable()

	extractTextSection(linkedExecutable)
}

func linkExecutable() string {
	/*
		Using ld linker link all object files
	*/

	// TODO:
	// cli.LogInfof("Linking PIC core + modules (%s)", strings.Join(modules, ","))

	outputPeFile := filepath.Join(ctx.LinkPIC.OutputPath, "linking", "payload.exe")
	linkerMapFile := filepath.Join(ctx.LinkPIC.OutputPath, "linking", "payload.linker.map")
	linkerScriptFile := filepath.Join(ctx.LinkPIC.OutputPath, "linking", "linker.ld")
	fs.MustCreateDirTree(linkerScriptFile)

	fs.MustWriteFile(linkerScriptFile, linkerScriptContent)

	objectFiles := getObjectFiles()

	if ctx.Debug {
		cli.LogDbg("Linking:")
		for _, f := range objectFiles {
			cli.LogDbgf(" - Linking: %s", f)
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

	output := utils.MingwLd(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("PIC-PE linked!")

	return outputPeFile
}

func getObjectFiles() []string {
	var objectFiles []string

	// Collecting core
	corePath := filepath.Join(ctx.LinkPIC.ObjectsPath, "core")
	for _, f := range fs.GetFilesByExtension(corePath, ".o") {
		objectFiles = append(objectFiles, f.FullPath)
	}

	// Collecting modules
	modulesPath := filepath.Join(ctx.LinkPIC.ObjectsPath, "modules")
	for _, module := range fs.GetChildDirs(modulesPath) {
		if !slices.Contains(ctx.LinkPIC.Modules, module) {
			continue
		}

		path := filepath.Join(ctx.LinkPIC.ObjectsPath, module)
		for _, f := range fs.GetFilesByExtension(path, ".o") {
			objectFiles = append(objectFiles, f.FullPath)
		}
	}

	return objectFiles
}

func extractTextSection(file string) {
	/*
		Using objcopy tool extract .text section from compiled executable.
	*/
	cli.LogInfo("Extracting '.text' section from PIC-PE...")
	outputFile := filepath.Join(ctx.LinkPIC.OutputPath, "payload.bin")

	output := utils.MingwObjcopy("-O", "binary", "--only-section=.text", file, outputFile)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("PIC payload extracted!")
}
