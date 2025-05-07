package lazypost

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
}

func New() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Batch()
}

func (m *Model) View() string {
	return ""
}
