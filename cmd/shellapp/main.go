package main

import (
	"cyoa"
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var story cyoa.Story

type model struct {
	chapter    string
	paragraphs []string
	choices    []cyoa.Option
	cursor     int
}

func initialModel() model {
	return model{
		chapter:    "intro",
		paragraphs: story["intro"].Paragraphs,
		choices:    story["intro"].Options,
		cursor:     0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
			break
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
			break
		case "enter":
			// TODO
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m model) View() string {
	s := "\n"
	for _, paragraph := range m.paragraphs {
		s += fmt.Sprintln(paragraph)
	}
	s += "\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s%s\n", cursor, choice.Text)
	}
	return s
}

func main() {
	filename := flag.String("file", "story.json", "file containing cyoa")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	content, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	story = content

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

/*

 */
