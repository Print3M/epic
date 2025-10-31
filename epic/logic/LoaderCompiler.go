package logic

import (
	_ "embed"
	"epic/cli"
	"epic/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type LoaderCompiler struct {
	PayloadPath string
	OutputPath  string
}

//go:embed assets/loader.c
var loaderContent string

func (lc *LoaderCompiler) ValidatePayloadPath() error {
	if !utils.PathExists(lc.PayloadPath) {
		return fmt.Errorf("payload file doesn't exist: %s", lc.PayloadPath)
	}

	if utils.MustIsDir(lc.PayloadPath) {
		return fmt.Errorf("payload path must be a file: %s", lc.PayloadPath)
	}

	return nil
}

func (lc *LoaderCompiler) ValidateOutputPath() error {
	if !utils.PathExists(lc.OutputPath) {
		return fmt.Errorf("output path doesn't exist: %s", lc.OutputPath)
	}

	if !utils.MustIsDir(lc.OutputPath) {
		return fmt.Errorf("output path must be a directory: %s", lc.OutputPath)
	}

	return nil
}

func (lc *LoaderCompiler) Run() {
	sourceFile := lc.createLoaderSourceFile()

	fmt.Println()
	lc.compileLoader(sourceFile)
}

func (lc *LoaderCompiler) createLoaderSourceFile() string {
	cli.LogInfof("Converting %s to bytes array", lc.PayloadPath)
	cStr, err := convertBinaryToCString(lc.PayloadPath)
	if err != nil {
		log.Fatalln(err)
	}

	assetsDir := filepath.Join(lc.OutputPath, "assets")
	utils.MustCreateDirTree(assetsDir)

	loaderWithPayload := strings.Replace(loaderContent, ":PAYLOAD:", cStr, 1)
	loaderFile := filepath.Join(assetsDir, "loader.c")
	utils.MustWriteFile(loaderFile, loaderWithPayload)

	cli.LogOkf("Payload injected -> %s", loaderFile)

	return loaderFile
}

func (lc *LoaderCompiler) compileLoader(sourceFile string) {
	outputFile := filepath.Join(lc.OutputPath, "loader.exe")

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
