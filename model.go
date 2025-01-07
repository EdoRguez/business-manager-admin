package main

import (
	"github.com/EdoRguez/business-manager-admin/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
)

type Model struct {
	Choice   int
	Chosen   bool
	Ticks    int
	Frames   int
	Progress float64
	IsLoaded bool
	Quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
	// return tick()
}

// Main update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !m.Chosen {
		return updateChoices(msg, m)
	}
	return updateChosen(msg, m)
}

// The main view, which just calls the appropriate sub-view
func (m Model) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if !m.Chosen {
		v := views.NewSelectViewModel(m.Choice)
		s = v.SelectView()
	} else {
		v := views.NewEndProgramModal(m.Choice, m.IsLoaded, m.Progress, m.Ticks)
		s = v.EndProgramView()
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

// Sub-update functions

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
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

	case tickMsg:
		if m.Ticks == 0 {
			m.Quitting = true
			return m, tea.Quit
		}
		m.Ticks--
		return m, tick()
	}

	return m, nil
}

// Update loop for the second view after a choice has been made
func updateChosen(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		if !m.IsLoaded {
			m.Frames++
			m.Progress = ease.OutBounce(float64(m.Frames) / float64(100))
			if m.Progress >= 1 {
				m.Progress = 1
				m.IsLoaded = true
				m.Ticks = 3
				return m, tick()
			}
			return m, frame()
		}

	case tickMsg:
		if m.IsLoaded {
			if m.Ticks == 0 {
				m.Quitting = true
				return m, tea.Quit
			}
			m.Ticks--
			return m, tick()
		}
	}

	return m, nil
}
