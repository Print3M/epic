package cli

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var _epic = `
@@@@@@@@  @@@@@@@   @@@   @@@@@@@    
@@@@@@@@  @@@@@@@@  @@@  @@@@@@@@
@@!       @@!  @@@  @@!  !@@     
!@!       !@!  @!@  !@!  !@!     
@!!!:!    @!@@!@!   !!@  !@!     
!!!!!:    !!@!!!    !!!  !!!     
!!:       !!:       !!:  :!!     
:!:       :!:       :!:  :!:     
:: ::::   ::        ::   ::: :::
: :: ::    :        :     :: :: :`

var _metadataTitle = `Extensible Position
Independent Code
`

var _metadataText = `Documentation:
‣ github.com/Print3M/epic

By Print3M
‣ print3m.github.io
`

func PrintBanner() {
	styleEpic := lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Foreground(lipgloss.Color("196")).
		Padding(0, 2, 2, 5)

	styleMetadataText := lipgloss.NewStyle()

	styleMetadataTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).Italic(true)

	metadataTitle := styleMetadataTitle.Render(_metadataTitle)
	metadataText := styleMetadataText.Render(_metadataText)
	metadata := lipgloss.JoinVertical(lipgloss.Left, metadataTitle, metadataText)

	banner := lipgloss.JoinHorizontal(
		lipgloss.Center,
		styleEpic.Render(_epic),
		metadata,
	)

	/*
	   	str := `
	       @@@@@@@@  @@@@@@@   @@@   @@@@@@@
	       @@@@@@@@  @@@@@@@@  @@@  @@@@@@@@       Extensible Position
	       @@!       @@!  @@@  @@!  !@@            Independent Code
	       !@!       !@!  @!@  !@!  !@!
	       @!!!:!    @!@@!@!   !!@  !@!            Documentation:
	       !!!!!:    !!@!!!    !!!  !!!            ‣ github.com/Print3M/epic
	       !!:       !!:       !!:  :!!
	       :!:       :!:       :!:  :!:            By Print3M
	        :: ::::   ::        ::   ::: :::       ‣ print3m.github.io
	       : :: ::    :        :     :: :: :
	   `
	*/
	fmt.Print(banner)
	fmt.Println()
}
