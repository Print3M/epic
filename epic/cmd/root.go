package cmd

import (
	"epic/cli"
	"epic/ctx"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "epic",
	Short:   "EPIC (Extensible Position Independent Code)",
	Long:    `EPIC is a CLI tool for automating modular PIC implant development and building process.`,
	Version: ctx.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ctx.NoColor {
			cli.DisableColors()
		}

		if ctx.Debug {
			cli.LogDbg("Debug mode enabled")

			if ctx.NoColor {
				cli.LogDbg("Colors disabled")
			}
		}
	},
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Flags available to all commands
	rootCmd.PersistentFlags().BoolVar(&ctx.Debug, "debug", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVar(&ctx.NoColor, "no-color", false, "disable colors output")
	rootCmd.PersistentFlags().BoolVar(&ctx.NoBanner, "no-banner", false, "disable epic banner")
	rootCmd.PersistentFlags().StringVar(&ctx.MingwGccPath, "mingw-w64-gcc", "", "path to MinGW-w64 GCC")
	rootCmd.PersistentFlags().StringVar(&ctx.MingwLdPath, "mingw-w64-ld", "", "path to MinGW-w64 ld")
	rootCmd.PersistentFlags().StringVar(&ctx.MingwObjcopyPath, "mingw-w64-objcopy", "", "path to MinGW-w64 objcopy")

	rootCmd.SetVersionTemplate(`{{printf "epic version %s\n" .Version}}`)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
