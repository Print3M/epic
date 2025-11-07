package cmd

import (
	"epic/cli"
	"epic/ctx"
	"epic/logic"
	"epic/utils"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var __plModules string

var pl logic.PICLinker

var linkCmd = &cobra.Command{
	Use:   "pic-link <path>",
	Short: "Link object files into PIC payload",
	Long:  `Link command links compiled object files (core + modules) into fully PIC payload.`,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		pl.ObjectsPath = args[0]

		if __plModules != "" {
			pl.Modules = utils.StringToSlice(__plModules, ",")
		}

		if err := pl.ValidateObjectsPath(); err != nil {
			return err
		}

		if err := pl.ValidateModules(); err != nil {
			return err
		}

		if err := pl.ValidateOutputPath(); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !ctx.NoBanner {
			cli.PrintBanner()
		}

		if ctx.Debug {
			cli.LogDbgf("Objects path: %s", pl.ObjectsPath)
			cli.LogDbgf("Output path: %s", pl.OutputPath)

			if pl.AllModules {
				cli.LogDbg("All modules included")
			} else {
				cli.LogDbgf("Modules: %s", strings.Join(pl.Modules, ","))
			}

			if ctx.MingwLdPath != "" {
				cli.LogDbgf("MinGW-w64 ld: %s", ctx.MingwLdPath)
			}

			if ctx.MingwLdPath != "" {
				cli.LogDbgf("MinGW-w64 objcopy: %s", ctx.MingwObjcopyPath)
			}
		}

		pl.Run()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	linkCmd.Flags().BoolVarP(&pl.AllModules, "all-modules", "a", false, "link all modules (ignore -m flag)")
	linkCmd.Flags().StringVarP(&__plModules, "modules", "m", "", "comma-separated list of modules")
	linkCmd.Flags().StringVarP(&pl.OutputPath, "output", "o", "", "output path (required)")

	// Mark required flags
	if err := linkCmd.MarkFlagRequired("output"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking flag as required: %v\n", err)
	}
}
