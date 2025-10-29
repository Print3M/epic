package pic

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
	"epic/utils"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

//go:embed assets/linker.ld
var linkerScriptContent string

func LinkPIC() {
	linkedExecutable := linkExecutable()

	fmt.Println()
	extractTextSection(linkedExecutable)
}

func linkExecutable() string {
	/*
		Using ld linker link all object files
	*/
	if modules := getLinkedModules(); len(modules) == 0 {
		cli.LogInfo("Linking PIC core (no modules)")

	} else {
		cli.LogInfof("Linking PIC core + modules (%s)", strings.Join(modules, ","))
	}

	assetsDir := filepath.Join(ctx.LinkPIC.OutputPath, "assets")
	utils.MustCreateDirTree(assetsDir)

	outputExecutable := filepath.Join(assetsDir, "payload.exe")
	linkerMapFile := filepath.Join(assetsDir, "payload.linker.map")
	linkerScriptFile := filepath.Join(assetsDir, "linker.ld")

	utils.MustWriteFile(linkerScriptFile, linkerScriptContent)

	objectFiles := getObjectFiles()

	for _, f := range objectFiles {
		cli.LogInfof(" ‣ %s ", f)

	}

	params := []string{
		"--gc-sections",
		"-Map", linkerMapFile,
		"-T", linkerScriptFile,
		"-o", outputExecutable,
		"--image-base=0x00",
	}

	if ctx.Debug {
		params = append(params, "--print-gc-sections")
	}

	params = append(params, objectFiles...)

	output := utils.MingwLd(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("PIC linked -> %s", outputExecutable)
	cli.LogOkf("Linking artifacts saved -> %s", assetsDir)

	return outputExecutable
}

func getObjectFiles() []string {
	var objectFiles []string

	// Collecting core
	corePath := filepath.Join(ctx.LinkPIC.ObjectsPath, "core")
	for _, f := range utils.GetFilesByExtensions(corePath, []string{".o"}) {
		objectFiles = append(objectFiles, f.FullPath)
	}

	// Collecting modules
	for _, module := range getLinkedModules() {
		path := filepath.Join(ctx.LinkPIC.ObjectsPath, "modules", module)

		for _, f := range utils.GetFilesByExtensions(path, []string{".o"}) {
			objectFiles = append(objectFiles, f.FullPath)
		}
	}

	return objectFiles
}

func getLinkedModules() []string {
	var modules []string
	path := filepath.Join(ctx.LinkPIC.ObjectsPath, "modules")

	for _, m := range utils.GetChildDirs(path) {
		if slices.Contains(ctx.LinkPIC.Modules, m) {
			modules = append(modules, m)
		}
	}

	return modules
}

func extractTextSection(file string) {
	/*
		Using objcopy tool extract .text section from compiled executable.
	*/
	cli.LogInfof("Extracting '.text' section from %s", file)
	outputFile := filepath.Join(ctx.LinkPIC.OutputPath, "payload.bin")

	output := utils.MingwObjcopy("-O", "binary", "--only-section=.text", file, outputFile)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("PIC payload extracted -> %s", outputFile)
}
