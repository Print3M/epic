package builder

import (
	"epic/ctx"
	"epic/fs"
	"epic/shell"
	"fmt"
)

func BuildStandalone() {
	fmt.Println("[*] Building standalone...")

	var sourceFiles []string

	for _, source := range fs.GetFilesByExtension(ctx.ProjectPath, ".c") {
		sourceFiles = append(sourceFiles, source.FullPath)
	}

	outputFile := fs.OutputPath("standalone.exe")
	params := []string{
		"-o", outputFile,
		"-w", "-Os",
		"-Wl,--subsystem,console,--entry,__main_pic",
		"-nostartfiles",
		"-static",
		"-s",
	}
	params = append(params, sourceFiles...)

	output := shell.MustExecuteProgram(ctx.CompilerPath, params...)

	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[*] Standalone built!")
}
