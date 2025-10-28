package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/monolith"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var monolithCmd = &cobra.Command{
	Use:   "monolith <path>",
	Short: "Build monolithic executable from project",
	Long:  `Monolith command builds a single monolithic executable from the specified project directory.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx.Monolith.ProjectPath = args[0]

		// Validate that path exists
		fileInfo, err := os.Stat(ctx.Monolith.ProjectPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("project directory does not exist: %s", ctx.Monolith.ProjectPath)
		}

		// Validate that it's a directory, not a file
		if !fileInfo.IsDir() {
			return fmt.Errorf("path must be a directory: %s", ctx.Monolith.ProjectPath)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx.Debug {
			cli.LogDbgf("Building monolith: %s", ctx.Monolith.ProjectPath)
			if ctx.Monolith.OutputPath != "" {
				cli.LogDbgf("Output: %s", ctx.Monolith.OutputPath)
			}
			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		// Your monolith build logic here
		monolith.CompileMonolith()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(monolithCmd)

	monolithCmd.Flags().StringVarP(&ctx.Monolith.OutputPath, "output", "o", "", "output path for generated executable")
	monolithCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")
}
