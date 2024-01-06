package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textarea textarea.Model
	value string
	err error
}

func initModel() model {
	t := textarea.New()
	t.Placeholder = "What's on your mind?"
	t.Focus()

	return model{
		textarea: t,
		value: "",
		err: nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			m.value = m.textarea.Value()
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.textarea.View()
}

func runEditor() (string, error) {
	m, err := tea.NewProgram(initModel()).Run()
	if err != nil {
		return "", err
	}
	e := m.(model)
	return e.value, e.err
}
