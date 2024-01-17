package editor

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textarea textarea.Model
	value    string
	err      error
	done     bool
	saved    bool
	savedAt  time.Time
}

func initModel() model {
	t := textarea.New()
	t.Placeholder = "What's on your mind?"
	t.Focus()

	return model{
		textarea: t,
		value:    "",
		err:      nil,
		done:     false,
		saved:    false,
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
		case tea.KeyCtrlS:
			m.value = m.textarea.Value()
			m.saved = true
			m.savedAt = time.Now()
		case tea.KeyEsc:
			m.done = true
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
	if m.done {
		// clear screen
		return ""
	}

	var sb strings.Builder
	sb.WriteString(m.textarea.View())
	sb.WriteString("\npress ctrl-s to save, esc to quit")

	if m.saved {
		sb.WriteRune('\n')
		sb.WriteString(fmt.Sprintf("last saved at %s", m.savedAt.Format(time.TimeOnly)))
	}

	return sb.String()
}

func Run() (string, error) {
	m, err := tea.NewProgram(initModel()).Run()
	if err != nil {
		return "", err
	}
	e := m.(model)
	return e.value, e.err
}
