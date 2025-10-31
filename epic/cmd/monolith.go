package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/logic"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mc logic.MonolithCompiler

var monolithCmd = &cobra.Command{
	Use:   "monolith <path>",
	Short: "Build monolithic executable from project",
	Long:  `Monolith command builds a single monolithic executable from the project directory.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		mc.ProjectPath = args[0]

		if err := mc.ValidateProjectPath(); err != nil {
			return err
		}

		if err := mc.ValidateOutputPath(); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !ctx.NoBanner {
			cli.PrintBanner()
		}

		if ctx.Debug {
			cli.LogDbgf("Project path: %s", mc.ProjectPath)
			cli.LogDbgf("Output path: %s", mc.OutputPath)

			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		mc.Run()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(monolithCmd)

	monolithCmd.Flags().StringVarP(&mc.OutputPath, "output", "o", "", "output path for generated executable (required)")

	// Mark required flags
	if err := monolithCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
