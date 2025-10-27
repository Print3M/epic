package builder

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
	"epic/fs"
	"epic/shell"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed assets/loader.c
var loaderContent string

func BuildLoader() {
	cli.LogInfo("Building loader...")

	var (
		loaderFile = createLoaderFile()
		outputFile = fs.OutputPath("loader.exe")
	)

	params := []string{
		"--sysroot", ctx.OutputPath,
		"-o", outputFile,
		"-static",
		"-s",
		loaderFile,
	}

	output := shell.MustExecuteProgram(ctx.CompilerPath, params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("Loader built!")
}

func createLoaderFile() string {
	cStr, err := binaryToCString(fs.OutputPath("payload.bin"))
	if err != nil {
		log.Fatalln(err)
	}

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := fs.OutputPath("assets", "loader.c")
	fs.MustCreateDirTree(loaderFile)
	fs.MustWriteFile(loaderFile, loaderWithPayload)

	cli.LogInfo("PIC payload injected into loader template")

	return loaderFile
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
