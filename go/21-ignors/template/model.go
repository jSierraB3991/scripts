package template

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Template struct {
	List     list.Model
	Choice   string
	Quitting bool
}

func (t Template) Init() tea.Cmd {
	return nil
}

func (t Template) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.List.SetWidth(msg.Width)
		return t, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			t.Quitting = true
			return t, tea.Quit
		case "enter":
			i, ok := t.List.SelectedItem().(Item)
			if ok {
				t.Choice = string(i)
			}
			return t, tea.Quit
		}
	}
	var cmd tea.Cmd
	t.List, cmd = t.List.Update(msg)
	return t, cmd
}

func (t Template) View() string {
	if t.Choice != "" {
		DownloadIgnor(t.Choice)
		return quitTextStyle.Render(fmt.Sprintf("ignor for %s download.", t.Choice))
	}
	if t.Quitting {
		return quitTextStyle.Render("Exitng ...")
	}
	return "\n" + t.List.View()
}
