package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/pic"
	"epic/utils"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var compileModules string

var compileCmd = &cobra.Command{
	Use:   "pic-compile <path>",
	Short: "Compile PIC project into object files",
	Long:  `Compile command compiles source code from the project directory and generates object files.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx.CompilePIC.ProjectPath = args[0]

		if !utils.PathExists(ctx.CompilePIC.ProjectPath) {
			return fmt.Errorf("project path does not exist: %s", ctx.CompilePIC.ProjectPath)
		}

		utils.ValidateProjectStructure(ctx.CompilePIC.ProjectPath)

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

		if compileModules != "" {
			ctx.CompilePIC.Modules = utils.StringToSlice(compileModules, ",")
		}

		pic.CompilePIC()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	compileCmd.Flags().StringVarP(&compileModules, "modules", "m", "", "comma-separated list of modules")
	compileCmd.Flags().StringVarP(&ctx.CompilePIC.OutputPath, "output", "o", "", "output path (required)")
	compileCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")

	// Mark required flags
	if err := compileCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
