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
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"log"
	"net/http"
)

var docStyle = lipgloss.NewStyle().Margin(1, 0)

type animeModel struct {
	animeList list.Model
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
	d.Styles.SelectedTitle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#EF6461"))
	d.Styles.NormalTitle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("192")).PaddingLeft(3)

	l := list.New(items, d, 0, 0)
	l.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#cacccf")).Background(lipgloss.Color("#EF6461"))

	l.SetShowStatusBar(false)

	m := animeModel{animeList: l, loading: true}
	m.animeList.Title = "Search results for Bleach"
	return &m
}

func (e errMsg) Error() string { return e.err.Error() }

func fetchSearchResults() tea.Msg {
	name := "bleach"
	animeList, err := searchAnime(name)
	if err != nil {
		return errMsg{err}
	}
	log.Printf("about to return Results")
	return result(animeList)
}

func searchAnime(name string) ([]animes, error) {
	var fullUrl string
	fullUrl = fmt.Sprintf("http://localhost:3000/anime/gogoanime/%s", name)

	resp, err := fetchJsonData(fullUrl)

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

func fetchJsonData(fullUrl string) ([]byte, error) {

	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Errorf("Failed to make the request: %v", err)
	}
	//defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Failed to read the response body: %v", err)
	}
	return body, nil
}

func (a animes) Title() string       { return a.title }
func (a animes) Description() string { return "" }
func (a animes) FilterValue() string { return a.title }

func (m animeModel) Init() tea.Cmd {
	m.loading = true
	return fetchSearchResults
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
	// if m.loading {
	// 	return docStyle.Render("Fetching Search Results...")
	// }

	return docStyle.Render(m.animeList.View())
}

func AnimeListMain() {
	m := NewResultsModel()

	f, err1 := tea.LogToFile("debug.log", "debug")
	if err1 != nil {
		log.Fatal("err: %w", err1)
	}
	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()

	if err != nil {
		log.Fatal(err)
	}
}
