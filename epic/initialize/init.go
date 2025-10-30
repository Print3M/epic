package initialize

import (
	"embed"
	"epic/cli"
	"epic/utils"
)

//go:embed all:template
var templateDir embed.FS

func InitProject(outputPath string) {
	utils.ExtractEmbeddedDir(templateDir, "template", outputPath)

	cli.LogOkf("Project structure created -> %s", outputPath)
}
