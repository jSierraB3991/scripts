package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "strings"
    "github.com/charmbracelet/bubbles/spinner"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    textinput textinput.Model
    spinner spinner.Model
    metaWeather *Client

    typing bool
    loading bool
    err error
    location Location
}

func (m *Model) Init() tea.Cmd {
    return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg: 
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "enter":
            if m.typing {
                query := strings.TrimSpace(m.textinput.Value())
                if query != "" {
                    m.typing = false
                    m.loading = true
                    return m, tea.Batch (
                        spinner.Tick,
                        m.featcWeather(query),
                    )
                }
            }
        case "esc":
            if !m.typing && !m.loading {
                m.typing = true
                m.err = nil
                return m, nil
            }
        }

    case GotWeather:
        m.loading = false;
        if msg.Err != nil {
            m.err = msg.Err
            return m, nil
        }
        m.location = msg.Location
        return m, nil
    }

    if m.typing {
        var cmd tea.Cmd
        m.textinput, cmd = m.textinput.Update(msg)
        return m, cmd
    }

    if m.loading {
        var cmd tea.Cmd
        m.spinner, cmd = m.spinner.Update(msg)
        return m, cmd
    }

    return m, nil
}

func (m *Model) View() string {
    if m.typing {
        return fmt.Sprintf("Enter you Location: \n%s", m.textinput.View())
    }
    if m.loading {
        return fmt.Sprintf("%s fetching weather... please wait.", m.spinner.View())
    }

    if m.err != nil {
        fmt.Sprintf("Could not fetch weather %v\n", m.err)
    }

    return fmt.Sprintf("Current Weather in %s is %.0f C", m.location.Title, m.location.ConsolidatedWeather[0].TheTemp)
}

type GotWeather struct {
    Err error
    Location Location
}

func (m Model) featcWeather(queryString string) tea.Cmd {
    return func() tea.Msg {
        location, err := m.metaWeather.LocationByQuery(context.Background(), queryString)
        if err != nil {
            return GotWeather{ Err: err }
        }
        return GotWeather {Location: location}
    }
}


func main() {
    t := textinput.NewModel()
    t.Focus()
    s := spinner.NewModel()
    s.Spinner = spinner.Dot


    err := tea.NewProgram(&Model{
        spinner: s,
        textinput: t,
        typing: true,
        metaWeather: &Client{ HTTPClient: http.DefaultClient },
    }).Start()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
