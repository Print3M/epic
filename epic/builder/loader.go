package builder

import (
	_ "embed"
	"epic/cli"
	"epic/fs"
	"epic/shell"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed assets/loader.c
var loaderContent string

func BuildLoader(flags *cli.CliFlags) {
	fmt.Println("[*] Building loader...")

	cStr, err := binaryToCString(filepath.Join(flags.OutputPath, "payload.bin"))
	if err != nil {
		log.Fatalln(err)
	}

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := filepath.Join(flags.OutputPath, "assets", "loader.c")
	fs.CreateDirTree(loaderFile)
	fs.MustWriteFile(loaderFile, loaderWithPayload)

	outputFile := filepath.Join(flags.OutputPath, "loader.exe")

	params := []string{
		"--sysroot", flags.OutputPath,
		"-o", outputFile,
		"-static",
		"-s",
		loaderFile,
	}

	output := shell.MustExecuteProgram(flags.CompilerPath, params...)

	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[*] Loader built!")
}

func binaryToCString(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	var result string
	for _, b := range data {
		result += fmt.Sprintf("\\x%02x", b)
	}

	return "" + result + "", nil
}
