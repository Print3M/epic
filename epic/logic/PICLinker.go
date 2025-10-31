package logic

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

type PICLinker struct {
	ObjectsPath string
	OutputPath  string
	Modules     []string
}

func (pl *PICLinker) ValidateObjectsPath() error {
	if !utils.PathExists(pl.ObjectsPath) {
		return fmt.Errorf("objects path does not exist: %s", pl.ObjectsPath)
	}

	if !utils.MustIsDir(pl.ObjectsPath) {
		return fmt.Errorf("objects path must be a directory: %s", pl.ObjectsPath)
	}

	if err := utils.ValidateProjectStructure(pl.ObjectsPath); err != nil {
		return err
	}

	return nil
}

func (pl *PICLinker) ValidateOutputPath() error {
	if !utils.PathExists(pl.OutputPath) {
		return fmt.Errorf("output path doesn't exist: %s", pl.OutputPath)
	}

	if !utils.MustIsDir(pl.OutputPath) {
		return fmt.Errorf("output path must be a directory: %s", pl.OutputPath)
	}

	return nil
}

func (pl *PICLinker) ValidateModules() error {
	// TODO: Validate modules

	return nil
}

func (pl *PICLinker) Run() {
	linkedExecutable := pl.linkExecutable()

	fmt.Println()
	pl.extractTextSection(linkedExecutable)
}

func (pl *PICLinker) linkExecutable() string {
	/*
		Using ld linker link all object files
	*/
	if modules := pl.getLinkedModules(); len(modules) == 0 {
		cli.LogInfo("Linking PIC core (no modules)")

	} else {
		cli.LogInfof("Linking PIC core + modules (%s)", strings.Join(modules, ","))
	}

	assetsDir := filepath.Join(pl.OutputPath, "assets")
	utils.MustCreateDirTree(assetsDir)

	outputExecutable := filepath.Join(assetsDir, "payload.exe")
	linkerMapFile := filepath.Join(assetsDir, "payload.linker.map")
	linkerScriptFile := filepath.Join(assetsDir, "linker.ld")

	utils.MustWriteFile(linkerScriptFile, linkerScriptContent)

	objectFiles := pl.getObjectFiles()

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

func (pl *PICLinker) getObjectFiles() []string {
	var objectFiles []string

	// Collecting core
	corePath := filepath.Join(pl.ObjectsPath, "core")
	for _, f := range utils.GetFilesByExtensions(corePath, []string{".o"}) {
		objectFiles = append(objectFiles, f.FullPath)
	}

	// Collecting modules
	for _, module := range pl.getLinkedModules() {
		path := filepath.Join(pl.ObjectsPath, "modules", module)

		for _, f := range utils.GetFilesByExtensions(path, []string{".o"}) {
			objectFiles = append(objectFiles, f.FullPath)
		}
	}

	return objectFiles
}

func (pl *PICLinker) getLinkedModules() []string {
	var modules []string
	path := filepath.Join(pl.ObjectsPath, "modules")

	for _, m := range utils.GetChildDirs(path) {
		if slices.Contains(pl.Modules, m) {
			modules = append(modules, m)
		}
	}

	return modules
}

func (pl *PICLinker) extractTextSection(file string) {
	/*
		Using objcopy tool extract .text section from compiled executable.
	*/
	cli.LogInfof("Extracting '.text' section from %s", file)
	outputFile := filepath.Join(pl.OutputPath, "payload.bin")

	output := utils.MingwObjcopy("-O", "binary", "--only-section=.text", file, outputFile)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("PIC payload extracted -> %s", outputFile)
}
