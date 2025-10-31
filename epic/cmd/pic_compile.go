package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/logic"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pc logic.PICCompiler

var compileCmd = &cobra.Command{
	Use:   "pic-compile <path>",
	Short: "Compile PIC project into object files",
	Long:  `Compile command compiles source code from the project directory and generates object files.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		pc.ProjectPath = args[0]

		if err := pc.ValidateProjectPath(); err != nil {
			return nil
		}

		if err := pc.ValidateOutputPath(); err != nil {
			return nil
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !ctx.NoBanner {
			cli.PrintBanner()
		}

		if ctx.Debug {
			cli.LogDbgf("Project path: %s", pc.ProjectPath)
			cli.LogDbgf("Output path: %s", pc.OutputPath)

			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		pc.Run()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	compileCmd.Flags().StringVarP(&pc.OutputPath, "output", "o", "", "output path (required)")
	// compileCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")

	// Mark required flags
	if err := compileCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
