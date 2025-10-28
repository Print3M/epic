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
	Short: "Link object files into standalone PIC payload",
	Long:  `Link command links compiled object files (core + modules) into standalone PIC payload.`,
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
		ctx.LinkPIC.ObjectsPath = args[0]

		if ctx.Debug {
			cli.LogDbgf("Linking from: %s\n", ctx.LinkPIC.ObjectsPath)

			if len(ctx.LinkPIC.Modules) > 0 {
				cli.LogDbgf("Modules: %s\n", strings.Join(ctx.LinkPIC.Modules, ","))
			}

			if ctx.MingwLdPath != "" {
				cli.LogDbgf("MinGW-w64 LD: %s\n", ctx.MingwLdPath)
			}

			if ctx.MingwLdPath != "" {
				cli.LogDbgf("MinGW-w64 objcopy: %s\n", ctx.MingwObjcopyPath)
			}
		}

		if linkModules != "" {
			ctx.LinkPIC.Modules = utils.StringToSlice(linkModules, ",")
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
