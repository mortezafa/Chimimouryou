package bubbleTeaUi

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
	titleText   lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("#EF6461")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).Width(80)
	s.titleText = lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("#B5AAAA"))
	return s
}

type searchModel struct {
	title       string
	width       int
	height      int
	output      string
	searchField textinput.Model
	styles      *Styles
}

type page int

func NewSearchModel() *searchModel {
	styles := DefaultStyles()
	searchField := textinput.New()
	searchField.Placeholder = "Ex: Chainsaw Man"
	searchField.Focus()
	return &searchModel{title: "Enter the anime you would like to search...", searchField: searchField, styles: styles}
}

// Defines the defaults state of the app
func (m searchModel) Init() tea.Cmd {
	return nil
}

// Function that re-renders our view with new state
func (m searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.output = m.searchField.Value()
			return m, nil
		case "left":
			m.searchField.SetValue("")
			return m, nil
		}

	}
	m.searchField, cmd = m.searchField.Update(msg)
	return m, cmd
}

// Function that renders our app
func (m searchModel) View() string {
	if m.width == 0 {
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.styles.titleText.Render(m.title),
			m.styles.InputField.Render(m.searchField.View())),
	)
}
