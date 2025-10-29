package cli

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func getTime() string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("black")).
		Background(lipgloss.Color("white")).
		Align(lipgloss.Center).
		Width(12)

	now := time.Now()

	return style.Render(now.Format("15:04:05"))
}

func LogErr(msg string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("124")).
		Align(lipgloss.Center).
		Width(10)

	fmt.Printf("%s%s %s", style.Render("ERR"), getTime(), msg)
	fmt.Println()
}

func LogErrf(format string, a ...any) {
	LogErr(fmt.Sprintf(format, a...))
}

func LogInfo(msg string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("32")).
		Align(lipgloss.Center).
		Width(10)

	fmt.Printf("%s%s %s", style.Render("INFO"), getTime(), msg)
	fmt.Println()
}

func LogInfof(format string, a ...any) {
	LogInfo(fmt.Sprintf(format, a...))
}

func LogOk(msg string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("34")).
		Align(lipgloss.Center).
		Width(10)

	fmt.Printf("%s%s %s", style.Render("OK"), getTime(), msg)
	fmt.Println()
}

func LogOkf(format string, a ...any) {
	LogOk(fmt.Sprintf(format, a...))
}

func LogDbg(msg string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("32")).
		Align(lipgloss.Center).
		Width(10)

	fmt.Printf("%s%s %s", style.Render("DBG"), getTime(), msg)
	fmt.Println()
}

func LogDbgf(format string, a ...any) {
	LogDbg(fmt.Sprintf(format, a...))
}
