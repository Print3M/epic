package logic

import (
	"embed"
	"epic/cli"
	"epic/utils"
	"fmt"
)

//go:embed all:init-template
var templateDir embed.FS

type ProjectInitializer struct {
	OutputPath string
}

func (pi *ProjectInitializer) ValidateOutputPath() error {
	if !utils.PathExists(pi.OutputPath) {
		return fmt.Errorf("output path does not exist: %s", pi.OutputPath)
	}

	if !utils.MustIsDir(pi.OutputPath) {
		return fmt.Errorf("output path must be a directory: %s", pi.OutputPath)
	}

	return nil
}

func (pi *ProjectInitializer) Run() {
	utils.ExtractEmbeddedDir(templateDir, "init-template", pi.OutputPath)

	cli.LogOkf("Project structure created -> %s", pi.OutputPath)
}
