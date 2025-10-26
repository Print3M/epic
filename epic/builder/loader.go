package builder

import (
	_ "embed"
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
	fmt.Println("[*] Building loader...")

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

	output := shell.MustExecuteProgram(ctx.MingwGccPath, params...)

	if len(output) > 0 {
		fmt.Println(output)
	}

	fmt.Println("[*] Loader built!")
}

func createLoaderFile() string {
	cStr, err := binaryToCString(fs.OutputPath("payload.bin"))
	if err != nil {
		log.Fatalln(err)
	}

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := fs.OutputPath("assets", "loader.c")
	fs.CreateDirTree(loaderFile)
	fs.MustWriteFile(loaderFile, loaderWithPayload)

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
