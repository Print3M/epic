package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/loader"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var loaderCmd = &cobra.Command{
	Use:   "loader <path>",
	Short: "Generate loader for binary payload",
	Long:  `Loader command generates a loader executable for the specified binary payload file.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx.Loader.PayloadPath = args[0]

		// Validate that path exists
		fileInfo, err := os.Stat(ctx.Loader.PayloadPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("payload file does not exist: %s", ctx.Loader.PayloadPath)
		}

		// Validate that it's a file, not a directory
		if fileInfo.IsDir() {
			return fmt.Errorf("path must be a file, not a directory: %s", ctx.Loader.PayloadPath)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx.Debug {
			cli.LogDbgf("Generating loader for payload: %s", ctx.Loader.PayloadPath)
			if ctx.Loader.OutputPath != "" {
				cli.LogDbgf("Output: %s", ctx.Loader.OutputPath)
			}
			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		// Your loader generation logic here
		loader.CompileLoader()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loaderCmd)

	loaderCmd.Flags().StringVarP(&ctx.Loader.OutputPath, "output", "o", "", "output path for generated loader")
	loaderCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")
}
