package monolith

import (
	"epic/cli"
	"epic/ctx"
	"epic/utils"
	"fmt"
)

func CompileMonolith() {
	/*
		Compile entire project into monolith (non-PIC) PE executable.
		It always compiles all modules.
	*/
	cli.LogInfo("Building monolith executable...")

	var sourceFiles []string

	for _, source := range utils.GetFilesByExtension(ctx.Monolith.ProjectPath, ".c") {
		sourceFiles = append(sourceFiles, source.FullPath)
	}

	params := []string{
		"-o", ctx.Monolith.OutputPath,
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

	cli.LogOk("Monolith executable built!")
}
