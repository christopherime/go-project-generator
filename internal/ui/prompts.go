package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbletea"
	"github.com/christopherime/go-project-generator/internal/generator"
	"github.com/christopherime/go-project-generator/internal/helpers"
)

type state int

const (
	enterName state = iota
	projectPathing
	selectType
)

type model struct {
	state       state
	projectPath string
	projectName string
	choices     []string
	selected    int
}

func (m model) Init() tea.Cmd {
	return nil
}

func StartPrompt() {
	p := tea.NewProgram(model{
		state:   projectPathing,
		choices: []string{"API", "Prometheus Exporter", "Simple Golang"},
	})

	_, err := p.Run()
	if err != nil {
		panic(err)
		return
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case projectPathing:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				m.state = enterName
				return m, nil
			case tea.KeyBackspace, tea.KeyCtrlH:
				if len(m.projectName) > 0 {
					m.projectName = m.projectName[:len(m.projectName)-1]
				}
			case tea.KeyRunes:
				m.projectName += string(msg.Runes)
			default:
				panic("Failed parsing project name")
			}
		}
	case enterName:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				m.state = selectType
				return m, nil
			case tea.KeyBackspace, tea.KeyCtrlH:
				if len(m.projectName) > 0 {
					m.projectName = m.projectName[:len(m.projectName)-1]
				}
			case tea.KeyRunes:
				m.projectName += string(msg.Runes)
			default:
				panic("Failed parsing project name")
			}
		}
	case selectType:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "up", "k":
				if m.selected > 0 {
					m.selected--
				}
			case "down", "j":
				if m.selected < len(m.choices)-1 {
					m.selected++
				}
			case "enter", " ":

				var project helpers.Project
				project.Type = m.choices[m.selected]
				project.Path = m.projectPath
				project.Name = m.projectName

				switch project.Type {

				case "API":
					err := generator.NewAPIProject(project)
					if err != nil {
						fmt.Printf("Error generating project: %v\n", err)
						return m, tea.Quit
					}
				case "Prometheus Exporter":
					err := generator.NewPrometheusProject(m.projectName)
					if err != nil {
						fmt.Printf("Error generating project: %v\n", err)
						return m, tea.Quit
					}
				case "Simple Golang":
					err := generator.NewGolangProject(m.projectName)
					if err != nil {
						fmt.Printf("Error generating project: %v\n", err)
						return m, tea.Quit
					}
				default:
					panic("unhandled choices")
				}
				return m, tea.Quit
			}
		}
	default:
		panic("unhandled tea model")
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case projectPathing:
		return fmt.Sprintf("Enter project path: %s\n", m.projectPath)
	case enterName:
		return fmt.Sprintf("Enter project name: %s\n", m.projectName)
	case selectType:
		s := "Select the type of project to generate:\n\n"
		for i, choice := range m.choices {
			cursor := " "
			if i == m.selected {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		return s + "\nPress enter to confirm."

	default:
		panic("unhandled tea view")
	}
	return ""
}
