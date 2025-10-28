package loader

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
	"epic/fs"
	"epic/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed assets/loader.c
var loaderContent string

func CompileLoader() {
	cli.LogInfo("Building loader...")

	var (
		loaderFile = createLoaderFile()
		outputFile = ctx.Loader.OutputPath
	)

	params := []string{
		// "--sysroot", ctx.OutputPath,
		"-o", outputFile,
		"-static",
		"-s",
		loaderFile,
	}

	output := utils.MingwGcc(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOk("Loader built!")
}

func createLoaderFile() string {
	cStr, err := binaryToCString(ctx.Loader.PayloadPath)
	if err != nil {
		log.Fatalln(err)
	}

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := filepath.Join(ctx.Loader.OutputPath, "assets", "loader.c")
	fs.MustCreateDirTree(loaderFile)
	fs.MustWriteFile(loaderFile, loaderWithPayload)

	cli.LogInfo("PIC payload injected into 'loader.c'")

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
