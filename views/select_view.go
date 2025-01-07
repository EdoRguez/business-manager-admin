package views

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	dotChar = " • "
)

const (
	ViewSelectView = iota
	ViewCreateCompany
	ViewEditCompany
	ViewEndProgram
)

var (
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
)

type SelectViewModel struct {
	selectOption int
}

func NewSelectViewModel(selectOption int) SelectViewModel {
	return SelectViewModel{
		selectOption: selectOption,
	}
}

// The first view, where you're choosing a task
func (m SelectViewModel) SelectView() string {
	tpl := "What to do today?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("esc, ctrl+c: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n",
		m.checkbox("Plant carrots", m.selectOption == ViewSelectView),
		m.checkbox("Go to the market", m.selectOption == ViewEndProgram),
	)

	return fmt.Sprintf(tpl, choices)
}

func (m SelectViewModel) checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
