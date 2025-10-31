package logic

import (
	"epic/cli"
	"epic/ctx"
	"epic/utils"
	"fmt"
	"path/filepath"
)

type MonolithCompiler struct {
	ProjectPath string
	OutputPath  string
}

func (mc *MonolithCompiler) ValidateProjectPath() error {
	if !utils.PathExists(mc.ProjectPath) {
		return fmt.Errorf("project path does not exist: %s", mc.ProjectPath)
	}

	if !utils.MustIsDir(mc.ProjectPath) {
		return fmt.Errorf("project path must be a directory: %s", mc.ProjectPath)
	}

	return nil
}

func (mc *MonolithCompiler) ValidateOutputPath() error {
	if !utils.PathExists(mc.OutputPath) {
		return fmt.Errorf("output path does not exist: %s", mc.OutputPath)
	}

	if !utils.MustIsDir(mc.OutputPath) {
		return fmt.Errorf("output path must be a directory: %s", mc.OutputPath)
	}

	return nil
}

func (mc *MonolithCompiler) Run() {
	/*
		Compile entire project into monolith (non-PIC) PE executable.
		It always compiles all modules.
	*/

	cli.LogInfof("Collecting source files from %s", mc.ProjectPath)

	var sourceFiles []string

	for _, source := range utils.GetFilesByExtensions(mc.ProjectPath, []string{".c", ".cpp"}) {
		sourceFiles = append(sourceFiles, source.FullPath)
	}

	if ctx.Debug {
		for _, f := range sourceFiles {
			cli.LogInfof(" ‣ %s ", f)
		}
	}

	cli.LogInfo("Compiling monolith executable...")

	outputFile := filepath.Join(mc.OutputPath, "monolith.exe")

	params := []string{
		"-o", outputFile,
		"-w", "-Os",
		"-Wl,--subsystem,console,--entry,__main_pic",
		"-static",
		"-ffixed-rbx",
		"-DMONOLITH",
		"-fno-lto",
		"-I", filepath.Join(mc.ProjectPath, "include"),
		"-I", mc.ProjectPath,
	}
	params = append(params, sourceFiles...)

	output := utils.MingwGcc(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("Monolith compiled -> %s", outputFile)
}
