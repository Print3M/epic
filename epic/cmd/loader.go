package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/logic"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lc logic.LoaderCompiler

var loaderCmd = &cobra.Command{
	Use:   "loader <path>",
	Short: "Build test loader with PIC payload",
	Long:  `Loader command builds a test loader executable with the specified payload.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		lc.PayloadPath = args[0]

		if err := lc.ValidatePayloadPath(); err != nil {
			return err
		}

		if err := lc.ValidateOutputPath(); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !ctx.NoBanner {
			cli.PrintBanner()
		}

		if ctx.Debug {
			cli.LogDbgf("Payload path: %s", lc.OutputPath)
			cli.LogDbgf("Output path: %s", lc.PayloadPath)

			if ctx.MingwGccPath != "" {
				cli.LogDbgf("MinGW-w64 GCC: %s", ctx.MingwGccPath)
			}
		}

		lc.Run()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loaderCmd)

	loaderCmd.Flags().StringVarP(&lc.OutputPath, "output", "o", "", "output path for generated loader (required)")
	// loaderCmd.Flags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")

	// Mark required flags
	if err := loaderCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
