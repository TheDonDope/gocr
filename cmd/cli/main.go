package main

import (
	"fmt"
	"os"

	"github.com/TheDonDope/gocr/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(config.InitialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
