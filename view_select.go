package main

import (
	"fmt"

	"github.com/EdoRguez/business-manager-admin/results"
	"github.com/EdoRguez/business-manager-admin/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar = " • "
)

var (
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
)

type ViewSelect struct {
	results        results.Results
	selectedOption int
}

func NewViewSelect(results results.Results) ViewSelect {
	return ViewSelect{
		results:        results,
		selectedOption: 0,
	}
}

// The first view, where you're choosing a task
func (v ViewSelect) SelectView() string {
	tpl := "What to do today?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc, ctrl+c: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n",
		v.checkbox("Create Company View", v.selectedOption == views.CreateCompanyView),
		v.checkbox("End Program View", v.selectedOption == views.EndProgramView),
	)

	return fmt.Sprintf(tpl, choices)
}

func (v ViewSelect) checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

// Update loop for the first view where you're choosing a task.
func (m Model) UpdateViewSelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			return m, frame()
		}
	}

	// case tickMsg:
	// 	if m.Ticks == 0 {
	// 		m.Quitting = true
	// 		return m, tea.Quit
	// 	}
	// 	m.Ticks--
	// 	return m, tick()
	// }

	return m, nil
}
