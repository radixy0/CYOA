package main

import (
	"cyoa"
	"flag"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
)

var story cyoa.Story

type model struct {
	chapter    string
	paragraphs []string
	choices    []cyoa.Option
	cursor     int
	width      int
	height     int
}

func initialModel() model {
	return model{
		chapter:    "intro",
		paragraphs: story["intro"].Paragraphs,
		choices:    story["intro"].Options,
		cursor:     0,
		width:      0,
		height:     0,
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
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			if len(m.choices) == 0 {
				return m, tea.Quit
			}
			choice := m.choices[m.cursor]
			nextChapter := choice.Chapter
			m.cursor = 0
			m.chapter = nextChapter
			m.choices = story[nextChapter].Options
			m.paragraphs = story[nextChapter].Paragraphs
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m model) View() string {
	s := "\n"
	maxWidth := 80
	displayWidth := min(maxWidth, m.width)
	for _, paragraph := range m.paragraphs {
		s += wordwrap.String(fmt.Sprintln(paragraph), displayWidth)
	}
	s += "\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s%s\n", cursor, strings.TrimLeft(indent.String(wordwrap.String(choice.Text, displayWidth-1), 1), " "))
	}

	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
