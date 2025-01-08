package main

import (
	"fmt"

	"github.com/EdoRguez/business-manager-admin/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar = " • "
)

var (
	checkboxStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	subtleStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	dotStyle               = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	possibleSelectecOption = 0
)

func ViewSelect(m Model) string {
	tpl := "What to do today?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc, ctrl+c: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n",
		checkbox("- Create Company", possibleSelectecOption == constants.View_Select),
		checkbox("- End Program", possibleSelectecOption == constants.View_EndProgram),
	)

	return fmt.Sprintf(tpl, choices)
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

// Update loop for the first view where you're choosing a task.
func UpdateViewSelect(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			possibleSelectecOption = 3
			// if m.Choice > 3 {
			// 	m.Choice = 3
			// }
		case "k", "up":
			possibleSelectecOption = 0
			// if m.Choice < 0 {
			// 	m.Choice = 0
			// }
		case "enter":
			possibleSelectecOption = 1
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
