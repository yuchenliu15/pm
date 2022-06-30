package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	path []string
}

func initialModel() model {
	path := []string{" "}
	return model {path: path}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m,  nil
}

func (m model) View() string {
	return "hi"
}

func main() {
	program := tea.NewProgram(initialModel())
	if err := program.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	} 
}


