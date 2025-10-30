package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/initialize"
	"epic/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	outputPath string
)

var initCmd = &cobra.Command{
	Use:   "init <path>",
	Short: "Create inital project structure",
	Long:  `Init command creates the initial project structure in the provided directory.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		outputPath = args[0]

		if !utils.PathExists(outputPath) {
			return fmt.Errorf("directory does not exist: %s", outputPath)
		}

		if !utils.MustIsDir(outputPath) {
			return fmt.Errorf("path must be a directory: %s", outputPath)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx.Debug {
			cli.LogDbgf("Output path: %s", outputPath)
		}

		initialize.InitProject(outputPath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
