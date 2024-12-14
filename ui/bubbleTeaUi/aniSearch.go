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
	output	string
	searchField textinput.Model 
	styles *Styles
	currentPage	page
	resmodel animeModel
}


type page int
const (
	 searchPage page = iota
	 resultsPage
)

func New(title string) *model {
	styles := DefaultStyles()
	searchField := textinput.New()
	searchField.Placeholder = "Ex: Chainsaw Man"
	searchField.Focus()
	return &model{title: title, searchField: searchField, styles: styles, currentPage: searchPage}
	
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
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.output = m.searchField.Value()
			log.Printf("searchValue: %s", m.output)
			m.currentPage = resultsPage
			return m, nil
		case "left":
			m.searchField.SetValue("")
			m.currentPage = searchPage;
			return m, nil
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

	switch m.currentPage {
	case searchPage:
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
	case resultsPage:
		log.Printf("Stuff: %s", m.resmodel.animeList.Title)
		return m.resmodel.View() 
	}
	return ""
}

func Main() {
	m := New("Enter the anime you would like to search...")

	f, err1 := tea.LogToFile("debug.log", "debug")
	if err1!= nil {
		log.Fatal("err: %w", err1)
	}
	defer f.Close()
	
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
