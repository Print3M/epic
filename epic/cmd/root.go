package cmd

import (
	"epic/cli"
	"epic/ctx"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "epic",
	Short:   "Epic - A powerful compilation and linking tool",
	Long:    `Epic is a CLI tool for compiling and linking with advanced options for module management.`,
	Version: ctx.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !ctx.NoBanner {
			printBanner()
		}

		if ctx.Debug {
			cli.LogDbg("Debug mode enabled")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Persistent flags available to all commands
	rootCmd.PersistentFlags().BoolVar(&ctx.Debug, "debug", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVar(&ctx.NoColor, "no-color", false, "disable colors output")
	rootCmd.PersistentFlags().BoolVar(&ctx.NoBanner, "no-banner", false, "disable epic banner")

	// Set custom version template
	rootCmd.SetVersionTemplate(`{{printf "epic version %s\n" .Version}}`)
}

func printBanner() {
	// TODO: Banner
	if ctx.NoColor {
		fmt.Println("=== EPIC ===")
	} else {
		// Add your colored banner here
		fmt.Println("\033[1;36m=== EPIC ===\033[0m")
	}
}

// Helper function to parse list of items
func parseList(str string, sep string) []string {
	items := strings.Split(str, sep)
	result := make([]string, 0, len(items))

	for _, item := range items {
		trimmed := strings.TrimSpace(item)

		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}
