package main

import (
	"fmt"
	"time"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

type model struct {
	spinner   spinner.Model
	done      bool
	files     []string
	msgdone   bool
	commitmsg string
}

type doneMsg struct{}

type fileFoundMsg struct {
	files []string
}

type commitmsg struct{}

func finishLoading() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		return doneMsg{}
	})
}

func generateCommitMessage() tea.Cmd {
	return tea.Tick(8*time.Second, func(t time.Time) tea.Msg {
		return commitmsg{}
	})
}

func findFiles() tea.Cmd {
	return func() tea.Msg {
		return fileFoundMsg{
			files: []string{
				"main.go",
				"README.md",
				"go.mod",
				"internal/git.go",
				"cmd/root.go",
			},
		}
	}
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Line
	return model{
		spinner: s,
	}
}
func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		finishLoading(),
		findFiles(),
		generateCommitMessage(),
	)

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case doneMsg:
		m.done = true
		return m, nil

	case fileFoundMsg:
		m.files = msg.files
		return m, nil

	case commitmsg:
		m.msgdone = true
		return m, nil
	}

	return m, nil
}

func (m model) View() tea.View {
	var content string

	if m.msgdone {
		content += "✓ Analyzed git repository\n\n"
		for _, file := range m.files {
			content += "  " + file + "\n"

		}
		content += "\n✓ Found 5 files\n\n"
		content += "Done!"

	} else if m.done {
		content += "✓ Analyzed git repository\n\n"
		for _, file := range m.files {
			content += "  " + file + "\n"

		}
		content += "\n✓ Found 5 files\n\n"
	} else {
		content += fmt.Sprintf(
			"%s Analyzing git repository...\n",
			m.spinner.View(),
		)
	}
	return tea.NewView(content)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
