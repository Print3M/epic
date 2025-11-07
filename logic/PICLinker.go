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
	AllModules  bool
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
	modules := utils.GetChildDirs(filepath.Join(pl.ObjectsPath, "modules"))

	for _, m := range pl.Modules {
		if !slices.Contains(modules, m) {
			return fmt.Errorf("unknown module: %s", m)
		}
	}

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
	if len(pl.Modules) == 0 {
		cli.LogInfo("Linking PIC core (no modules)")

	} else {
		cli.LogInfof("Linking PIC core + modules (%s)", strings.Join(pl.Modules, ","))
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
		"-T", linkerScriptFile,
		"-o", outputExecutable,
		"-Map", linkerMapFile,
		"--no-seh",
		"--gc-sections",
		"--image-base=0x00",
		"--file-alignment=4",
		"--section-alignment=4",
		"--disable-runtime-pseudo-reloc",
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
	objectFiles = append(objectFiles, pl.getModuleObjectFiles()...)

	if len(objectFiles) == 0 {
		cli.LogErrf("No object files (*.o) found in %s", pl.ObjectsPath)
	}

	return objectFiles
}

func (pl *PICLinker) getModuleObjectFiles() []string {
	var objectFiles []string

	if pl.AllModules {
		path := filepath.Join(pl.ObjectsPath, "modules")

		for _, f := range utils.GetFilesByExtensions(path, []string{".o"}) {
			objectFiles = append(objectFiles, f.FullPath)
		}
	} else {
		for _, module := range pl.Modules {
			path := filepath.Join(pl.ObjectsPath, "modules", module)

			for _, f := range utils.GetFilesByExtensions(path, []string{".o"}) {
				objectFiles = append(objectFiles, f.FullPath)
			}
		}
	}

	return objectFiles
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
