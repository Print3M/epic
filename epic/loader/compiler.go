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
	sourceFile := createLoaderSourceFile()

	fmt.Println()
	compileLoader(sourceFile)
}

func compileLoader(sourceFile string) {
	outputFile := filepath.Join(ctx.Loader.OutputPath, "loader.exe")

	params := []string{
		// "--sysroot", ctx.OutputPath,
		"-o", outputFile,
		"-static",
		"-s",
		sourceFile,
	}

	cli.LogInfof("Compiling %s", sourceFile)
	output := utils.MingwGcc(params...)
	if len(output) > 0 {
		fmt.Println(output)
	}

	cli.LogOkf("Loader compiled -> %s", outputFile)
}

func createLoaderSourceFile() string {
	cli.LogInfof("Converting %s to bytes array", ctx.Loader.PayloadPath)
	cStr, err := convertBinaryToCString(ctx.Loader.PayloadPath)
	if err != nil {
		log.Fatalln(err)
	}

	assetsDir := filepath.Join(ctx.Loader.OutputPath, "assets")
	utils.MustCreateDirTree(assetsDir)

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := filepath.Join(assetsDir, "loader.c")
	utils.MustWriteFile(loaderFile, loaderWithPayload)

	cli.LogOkf("Payload injected -> %s", loaderFile)

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
