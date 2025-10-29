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

var linkModules string

var linkCmd = &cobra.Command{
	Use:   "pic-link <path>",
	Short: "Link object files into PIC payload",
	Long:  `Link command links compiled object files (core + modules) into fully PIC payload.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx.LinkPIC.ObjectsPath = args[0]

		if !utils.PathExists(ctx.LinkPIC.ObjectsPath) {
			return fmt.Errorf("project path does not exist: %s", ctx.LinkPIC.ObjectsPath)
		}

		utils.ValidateProjectStructure(ctx.LinkPIC.ObjectsPath)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if linkModules != "" {
			ctx.LinkPIC.Modules = utils.StringToSlice(linkModules, ",")
		}

		if ctx.Debug {
			cli.LogDbgf("Objects path: %s", ctx.LinkPIC.ObjectsPath)
			cli.LogDbgf("Output path: %s", ctx.LinkPIC.OutputPath)
			cli.LogDbgf("Modules: %s", strings.Join(ctx.LinkPIC.Modules, ","))

			if ctx.MingwLdPath != "" {
				cli.LogDbgf("MinGW-w64 ld: %s", ctx.MingwLdPath)
			}

			if ctx.MingwLdPath != "" {
				cli.LogDbgf("MinGW-w64 objcopy: %s", ctx.MingwObjcopyPath)
			}
		}

		pic.LinkPIC()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	linkCmd.Flags().StringVarP(&linkModules, "modules", "m", "", "comma-separated list of modules")
	linkCmd.Flags().StringVar(&ctx.MingwLdPath, "mingw-w64-ld", "", "path to MinGW-w64 ld")
	linkCmd.Flags().StringVar(&ctx.MingwObjcopyPath, "mingw-w64-objcopy", "", "path to MinGW-w64 objcopy")
	linkCmd.Flags().StringVarP(&ctx.LinkPIC.OutputPath, "output", "o", "", "output path (required)")

	// Mark required flags
	if err := linkCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
