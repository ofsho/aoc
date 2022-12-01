package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"2022/day1"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			switch m.table.SelectedRow()[0] {
			case "1":
				fmt.Println(day1.Day()[0])
				fmt.Println(day1.Day()[1])
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
	columns := []table.Column{
		{Title: "Day", Width: 4},
		{Title: "Name", Width: 25},
		{Title: "Completed", Width: 4},
		{Title: "Language", Width: 8},
	}

	rows := []table.Row{
		{"1", "Elf Calories", "no", "Go"},
		{"2", "", "no", "Go"},
		{"3", "", "no", "Go"},
		{"4", "", "no", "Go"},
		{"5", "", "no", "Go"},
		{"6", "", "no", "Go"},
		{"7", "", "no", "Go"},
		{"8", "", "no", "Go"},
		{"9", "", "no", "Go"},
		{"10", "", "no", "Go"},
		{"11", "", "no", "Go"},
		{"12", "", "no", "Go"},
		{"13", "", "no", "Go"},
		{"14", "", "no", "Go"},
		{"15", "", "no", "Go"},
		{"16", "", "no", "Go"},
		{"17", "", "no", "Go"},
		{"18", "", "no", "Go"},
		{"19", "", "no", "Go"},
		{"20", "", "no", "Go"},
		{"21", "", "no", "Go"},
		{"22", "", "no", "Go"},
		{"23", "", "no", "Go"},
		{"24", "", "no", "Go"},
		{"25", "", "no", "Go"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
