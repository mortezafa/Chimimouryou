package bubbleTeaUi

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

var docStyle = lipgloss.NewStyle().Margin(1, 0)


type animeModel struct {
	animeList list.Model
	styles    lipgloss.Style
}

type animes struct {
	title string
}

func (a animes) Title() string  { return a.title }
func (a animes) Description() string  { return "" }
func (a animes) FilterValue() string  { return a.title }


func (m animeModel) Init() tea.Cmd {
	return nil	
}

func (m animeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.animeList.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.animeList, cmd = m.animeList.Update(msg)
	return m, cmd
}


func (m animeModel) View() string {
	return docStyle.Render(m.animeList.View())
	
}

func AnimeListMain() {
	items := []list.Item{
		animes{title: "Tokyo Ghoul"},
		animes{title: "Tokyo Ghoul:RE"},
		animes{title: "Tokyo Ghoul Root A"},
	}
	
	d := list.NewDefaultDelegate()
	d.ShowDescription = false
	d.Styles.SelectedTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("32")).PaddingLeft(3)
	d.Styles.NormalTitle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("192")).PaddingLeft(3)
	
	l := list.New(items, d, 0,0)
	l.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Background(lipgloss.NoColor{})
	
	l.SetShowStatusBar(false)	
	
	
	m := animeModel{animeList: l}
	m.animeList.Title = "Search results for Tokyo Ghoul"
	
	
	
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