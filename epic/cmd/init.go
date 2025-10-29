package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/monolith"
	"epic/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// TODO: Init project, drop entire project to disk
var initCmd = &cobra.Command{
	Use:   "monolith <path>",
	Short: "Build monolithic executable from project",
	Long:  `Monolith command builds a single monolithic executable from the project directory.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx.Monolith.ProjectPath = args[0]

		if !utils.PathExists(ctx.Monolith.ProjectPath) {
			return fmt.Errorf("project directory does not exist: %s", ctx.Monolith.ProjectPath)
		}

		if !utils.MustIsDir(ctx.Monolith.ProjectPath) {
			return fmt.Errorf("path must be a directory: %s", ctx.Monolith.ProjectPath)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx.Debug {
			cli.LogDbgf("Project path: %s", ctx.Monolith.ProjectPath)
			cli.LogDbgf("Output path: %s", ctx.Monolith.OutputPath)

			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		monolith.CompileMonolith()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
