package builder

import (
	"epic/cli"
	"epic/ctx"
	"epic/fs"
	"epic/shell"
	"fmt"
)

func BuildNonPIC() {
	cli.LogInfo("Building non-PIC executable...")

	var sourceFiles []string

	for _, source := range fs.GetFilesByExtension(ctx.ProjectPath, ".c") {
		sourceFiles = append(sourceFiles, source.FullPath)
	}

	outputFile := fs.OutputPath("non-pic.exe")
	params := []string{
		"-o", outputFile,
		"-w", "-Os",
		"-Wl,--subsystem,console,--entry,__main_pic",
		"-nostartfiles",
		"-static",
		"-ffixed-rbx",
		"-DNONPIC",
	}
	params = append(params, sourceFiles...)

	output := shell.MustExecuteProgram(ctx.CompilerPath, params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("Non-PIC executable built!")
}
