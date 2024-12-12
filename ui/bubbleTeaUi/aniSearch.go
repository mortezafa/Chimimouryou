package bubbleTeaUi

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

type Styles struct{
	BorderColor lipgloss.Color
	InputField lipgloss.Style
	titleText lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("#EF6461")
	s.InputField = lipgloss.NewStyle().BorderForeground(s.BorderColor).BorderStyle(lipgloss.RoundedBorder()).
		Padding(1).Width(80)
	s.titleText = lipgloss.NewStyle().Padding(1).Foreground(lipgloss.Color("#B5AAAA"))
	return s
}

type model struct {
	title string
	width int
	height int
	searchField textinput.Model 
	styles *Styles
}

func New(title string) *model {
	styles := DefaultStyles()
	searchField := textinput.New()
	searchField.Placeholder = "Ex: Chainsaw Man"
	searchField.Focus()
	return &model{title: title, searchField: searchField, styles: styles}
	
}

// Defines the defaults state of the app
func (m model) Init() tea.Cmd {
	return nil
}

// Function that re-renders our view with new state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd)  {
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
			return m, tea.Quit
		}
	}
	m.searchField, cmd = m.searchField.Update(msg)
	return m, cmd
}


// Function that renders our app
func (m model)View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
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

func RunTUI() (string, error) {

	title := "Type in the Anime you would like to search for..."

	m := New(title)
	p:= tea.NewProgram(m, tea.WithAltScreen())
	
	finalModel, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	output := finalModel.(model)
	return output.searchField.Value(), nil
	
}

func Main() {
	_, err := RunTUI()	
	if err != nil {
		log.Fatal(err)	
	}
	
}

