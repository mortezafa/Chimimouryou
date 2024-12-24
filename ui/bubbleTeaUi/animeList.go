/*
* Model manages state.
* Update is going to switch on our state, and send updates accordingly...?
*
*
*
*
*
*
 */
package bubbleTeaUi

import (
	"Chimimouryou/JsonsStrcuts"
	"Chimimouryou/utils"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

var docStyle = lipgloss.NewStyle().Margin(1, 0)

type animeModel struct {
	animeList list.Model
	searchTerm string
	animeID string
	selectedAnimeName string
	err       error
	loading   bool
}

type animes struct {
	title string
	id    string
	image string
}

type (
	result []animes
	errMsg struct{ err error }
)

func NewResultsModel() *animeModel {
	items := []list.Item{}

	d := list.NewDefaultDelegate()
	d.ShowDescription = false
	d.Styles.SelectedTitle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#e82017"))
	d.Styles.NormalTitle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("192")).PaddingLeft(3)

	l := list.New(items, d, 0, 0)
	l.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#cacccf")).Background(lipgloss.Color("#e82017"))

	l.SetShowStatusBar(false)

	m := animeModel{animeList: l, loading: true}
	m.animeList.Title = fmt.Sprintf("Search results for %s", m.searchTerm)
	return &m
}

func (e errMsg) Error() string { return e.err.Error() }

func fetchSearchResults(name string) tea.Cmd {
	return func() tea.Msg {
		animeList, err := searchAnime(name)
		if err != nil {
			return errMsg{err}
		}
		log.Printf("about to return Results")
		return result(animeList)
	}
}

func searchAnime(name string) ([]animes, error) {
	var fullUrl string
	fullUrl = fmt.Sprintf("http://localhost:3000/anime/gogoanime/%s", name)

	resp, err := utils.FetchJsonData(fullUrl)

	var animeSearchQuery JsonsStrcuts.AnimeSearchQuery
	err = json.Unmarshal(resp, &animeSearchQuery)
	if err != nil {
		fmt.Errorf("Failed to parse the response body: %v", err)
		return nil, nil
	}

	animeList := []animes{}
	for _, source := range animeSearchQuery.Results {
		animeList = append(animeList, animes{
			title: source.Title,
			id:    source.ID,
			image: source.Image,
		})
	}

	return animeList, nil

}

func (m animeModel) SetSearchTerm(term string) (animeModel, tea.Cmd) {
	m.loading = true
	m.searchTerm = term
	m.animeList.Title = fmt.Sprintf("Search Results for %s", m.searchTerm)
	return m, fetchSearchResults(term)
}



func (a animes) Title() string       { return a.title }
func (a animes) Description() string { return "" }
func (a animes) FilterValue() string { return a.title }

func (m animeModel) Init() tea.Cmd {
	m.loading = true
	return fetchSearchResults(m.searchTerm)
}

func (m animeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			m.animeID = m.animeList.SelectedItem().(animes).id
			m.selectedAnimeName = m.animeList.SelectedItem().(animes).title
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.animeList.SetSize(msg.Width-h, msg.Height-v)

	case result:
		m.loading = false

		searchResults := []animes(msg)
		items := []list.Item{}

		for _, anime := range searchResults {
			items = append(items, animes{
				title: anime.title,
				id:    anime.id,
				image: anime.id,
			})
		}

		log.Printf("Search Results: %s", items)
		m.animeList.SetItems(items)
		log.Println("Current Anime List when NewResults is called")
		return m, nil
	case errMsg:
		m.err = msg
	}

	var cmd tea.Cmd
	m.animeList, cmd = m.animeList.Update(msg)
	return m, cmd
}

func (m animeModel) View() string {

	log.Printf("poo %v", m.loading)
	if m.loading {
		return docStyle.Render("Fetching Search Results...")
	}

	return docStyle.Render(m.animeList.View())
}

