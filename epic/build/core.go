package build

import (
	"epic/cli"
	"epic/fs"
	"epic/shell"
	"fmt"
	"log"
	"path/filepath"
)

func BuildCore(flags *cli.CliFlags) {
	files := fs.GetAllFileEntriesFlat(filepath.Join(flags.InputPath, "core"))

	for _, file := range files {
		if !fs.IsSourceFile(file.Name) || file.IsDir {
			continue
		}

		objectFileName := fs.ReplaceExtension(file.Name, "o")

		params := []string{
			"-c",
			file.Path,
			"-o",
			filepath.Join(flags.OutputPath, objectFileName),
			"-nostdlib",
			"-fPIC",
			"-nostartfiles",
			"-Os",
			"-fno-asynchronous-unwind-tables",
			"-ffreestanding",
			"-fno-builtin",
			"-ffunction-sections",
			"-fno-ident",
			"-falign-jumps=1",
			"-mno-sse",
			"-mno-mmx",
			"-mgeneral-regs-only",
			"-nostdinc",
			"-mno-stack-arg-probe",
			"-mno-red-zone",
			"-fdiagnostics-color=always",
			"-std=c17",
		}

		fmt.Printf("[*] Building file: %s\n\n", file.Path)
		output, err := shell.ExecuteProgram(flags.CompilerPath, params...)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(output)
	}

}
