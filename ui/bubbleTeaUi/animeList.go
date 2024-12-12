package bubbleTeaUi

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type AnimeListStyles struct {
	listText	lipgloss.Style	
}

func aniStyles() *AnimeListStyles {
	s := new(AnimeListStyles)
	s.listText = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	return s
	
}

type animeModel struct {
	animeList list.Model
	styles    *AnimeListStyles
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
	return m.styles.listText.Render(docStyle.Render(m.animeList.View()))
	
}

func AnimeListMain() {
	items := []list.Item{
		animes{title: "Tokyo Ghoul"},
		animes{title: "Tokyo Ghoul:RE"},
		animes{title: "Tokyo Ghoul Root A"},
	}
	
	styles := aniStyles()
	l :=  list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(false)
	
	
	m := animeModel{animeList: l, styles: styles}
	m.animeList.Title = "Search results for Tokyo Ghoul"
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	
	if err != nil {
		log.Fatal(err)
	}
	
}