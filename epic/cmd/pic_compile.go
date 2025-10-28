package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/pic"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var compileModules string

var compileCmd = &cobra.Command{
	Use:   "pic-compile <path>",
	Short: "Compile source code with specified modules",
	Long:  `Compile command processes source code from the specified directory and generates output files.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		if ctx.CompilePIC.OutputPath == "" {
			return fmt.Errorf("--output/-o flag is required")
		}

		// Validate path exists
		ctx.CompilePIC.ProjectPath = args[0]
		if _, err := os.Stat(ctx.CompilePIC.ProjectPath); os.IsNotExist(err) {
			return fmt.Errorf("source path does not exist: %s", ctx.CompilePIC.ProjectPath)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx.Debug {
			cli.LogDbgf("Compiling from: %s", ctx.CompilePIC.ProjectPath)
			cli.LogDbgf("Output: %s", ctx.CompilePIC.OutputPath)

			if len(ctx.CompilePIC.Modules) > 0 {
				cli.LogDbgf("Modules: %s", strings.Join(ctx.CompilePIC.Modules, ","))
			}

			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		// Parse modules if provided
		if compileModules != "" {
			ctx.CompilePIC.Modules = parseList(compileModules, ",")
		}

		// Your compilation logic here
		pic.CompilePIC()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	compileCmd.Flags().StringVarP(&compileModules, "modules", "m", "", "comma-separated list of modules")
	compileCmd.Flags().StringVarP(&ctx.CompilePIC.OutputPath, "output", "o", "", "path to output directory (required)")
	compileCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")

	// Mark required flags
	if err := compileCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
