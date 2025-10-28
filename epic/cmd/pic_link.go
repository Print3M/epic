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

var linkModules string

var linkCmd = &cobra.Command{
	Use:   "pic-link <path>",
	Short: "Link compiled object files",
	Long:  `Link command processes compiled object files from the specified directory.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate path exists
		ctx.LinkPIC.ObjectsPath = args[0]
		if _, err := os.Stat(ctx.LinkPIC.ObjectsPath); os.IsNotExist(err) {
			return fmt.Errorf("source path does not exist: %s", ctx.LinkPIC.ObjectsPath)
		}

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
			ctx.LinkPIC.Modules = parseList(linkModules, ",")
		}

		// Your linking logic here
		pic.LinkPIC()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	linkCmd.Flags().StringVarP(&linkModules, "modules", "m", "", "comma-separated list of modules")
	linkCmd.Flags().StringVar(&ctx.MingwLdPath, "mingw-w64-ld", "", "path to MinGW-w64 LD")
	linkCmd.Flags().StringVar(&ctx.MingwObjcopyPath, "mingw-w64-objcopy", "", "path to MinGW-w64 objcopy")
}
