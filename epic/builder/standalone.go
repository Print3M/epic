package builder

import (
	"epic/cli"
	"epic/fs"
	"epic/shell"
	"fmt"
	"path/filepath"
)

func BuildStandalone(flags *cli.CliFlags) {
	fmt.Println("[*] Building standalone...")

	var sourceFiles []string

	for _, source := range fs.GetFilesByExtension(flags.InputPath, ".c") {
		sourceFiles = append(sourceFiles, source.FullPath)
	}

	outputFile := filepath.Join(flags.OutputPath, "standalone.exe")
	params := []string{
		"-o", outputFile,
		"-w", "-Os",
		"-Wl,--subsystem,console",
		"-static",
		"-s",
	}
	params = append(params, sourceFiles...)

	output := shell.MustExecuteProgram(flags.CompilerPath, params...)

	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[*] Standalone built!")
}
