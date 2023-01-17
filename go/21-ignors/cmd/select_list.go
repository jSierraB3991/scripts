package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jSierraB3991/scripts/21-ignors/template"
)

const listHeight = 14

var (
	itemStyle       = lipgloss.NewStyle().PaddingLeft(4)
	titleStyle      = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

func Run(templates []string) {

	var items []list.Item
	for _, templateLan := range templates {
		items = append(items, template.Item(templateLan))
	}

	const defaultWidth = 20

	l := list.New(items, template.ItemDelegate{}, defaultWidth, listHeight)
	l.Title = "What do you want for download ignor?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	t := template.Template{List: l}

	if _, err := tea.NewProgram(t).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
