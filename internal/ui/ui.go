package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jublx/lazypost/internal/ui/lazypost"
)

func NewProgram() *tea.Program {
	return tea.NewProgram(lazypost.New(), tea.WithAltScreen())
}
