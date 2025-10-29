package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/loader"
	"epic/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var loaderCmd = &cobra.Command{
	Use:   "loader <path>",
	Short: "Build test loader with PIC payload",
	Long:  `Loader command builds a test loader executable with the specified payload.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx.Loader.PayloadPath = args[0]

		if !utils.PathExists(ctx.Loader.PayloadPath) {
			return fmt.Errorf("payload file doesn't exist: %s", ctx.Loader.PayloadPath)
		}

		if utils.MustIsDir(ctx.Loader.PayloadPath) {
			return fmt.Errorf("payload path must be a file: %s", ctx.Loader.PayloadPath)
		}

		if !utils.PathExists(ctx.Loader.OutputPath) {
			return fmt.Errorf("output path doesn't exist: %s", ctx.Loader.OutputPath)
		}

		if !utils.MustIsDir(ctx.Loader.OutputPath) {
			return fmt.Errorf("output path must be a directory: %s", ctx.Loader.OutputPath)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx.Debug {
			cli.LogDbgf("Payload path: %s", ctx.Loader.PayloadPath)
			cli.LogDbgf("Output path: %s", ctx.Loader.OutputPath)

			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		loader.CompileLoader()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loaderCmd)

	loaderCmd.Flags().StringVarP(&ctx.Loader.OutputPath, "output", "o", "", "output path for generated loader (required)")
	loaderCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")

	// Mark required flags
	if err := loaderCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
