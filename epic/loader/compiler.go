package loader

import (
	_ "embed"
	"epic/cli"
	"epic/ctx"
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

	var (
		loaderFile = createLoaderFile()
		outputFile = filepath.Join(ctx.Loader.OutputPath, "loader.exe")
	)

	params := []string{
		// "--sysroot", ctx.OutputPath,
		"-o", outputFile,
		"-static",
		"-s",
		loaderFile,
	}

	cli.LogInfo("Compiling loader...")
	output := utils.MingwGcc(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("Loader built -> %s", outputFile)
}

func createLoaderFile() string {
	cStr, err := convertBinaryToCString(ctx.Loader.PayloadPath)
	if err != nil {
		log.Fatalln(err)
	}

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := filepath.Join(ctx.Loader.OutputPath, "assets", "loader.c")
	utils.MustCreateDirTree(loaderFile)
	utils.MustWriteFile(loaderFile, loaderWithPayload)

	cli.LogInfof("PIC payload injected -> %s", loaderFile)

	return loaderFile
}

func convertBinaryToCString(file string) (string, error) {
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
