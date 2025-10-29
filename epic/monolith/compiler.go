package monolith

import (
	"epic/cli"
	"epic/ctx"
	"epic/utils"
	"fmt"
	"path/filepath"
)

func CompileMonolith() {
	/*
		Compile entire project into monolith (non-PIC) PE executable.
		It always compiles all modules.
	*/

	cli.LogInfof("Collecting source files from %s", ctx.Monolith.ProjectPath)

	var sourceFiles []string

	for _, source := range utils.GetFilesByExtensions(ctx.Monolith.ProjectPath, []string{".c", ".cpp"}) {
		sourceFiles = append(sourceFiles, source.FullPath)
	}

	if ctx.Debug {
		for _, f := range sourceFiles {
			cli.LogInfof(" ‣ %s ", f)
		}
	}

	cli.LogInfo("Compiling monolith executable...")

	outputFile := filepath.Join(ctx.Monolith.OutputPath, "monolith.exe")

	params := []string{
		"-o", outputFile,
		"-w", "-Os",
		"-Wl,--subsystem,console,--entry,__main_pic",
		"-nostartfiles",
		"-static",
		"-ffixed-rbx",
		"-DMONOLITH",
	}
	params = append(params, sourceFiles...)

	output := utils.MingwGcc(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("Monolith compiled -> %s", outputFile)
}
