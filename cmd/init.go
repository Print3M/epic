package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/logic"

	"github.com/spf13/cobra"
)

var pi logic.ProjectInitializer

var initCmd = &cobra.Command{
	Use:   "init <path>",
	Short: "Create inital project structure",
	Long:  `Init command creates the initial project structure in the provided directory.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		pi.OutputPath = args[0]

		if err := pi.ValidateOutputPath(); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !ctx.NoBanner {
			cli.PrintBanner()
		}

		if ctx.Debug {
			cli.LogDbgf("Output path: %s", pi.OutputPath)
		}

		pi.Run()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
