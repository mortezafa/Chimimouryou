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

type episodeModel struct{
	episodeList list.Model
	animetitle string
	animeId    string
	err        error
	loading   bool
	episodeId string
}

type episodes struct{
	id string
	epNum int
}

type (
	epRes []episodes
	epErrMsg struct{ err error }
)

func (e episodes) Title() string       { return e.id }
func (e episodes) Description() string { return "" }
func (e episodes) FilterValue() string { return e.id }

func NewEpModel() *episodeModel {
	items := []list.Item{}

	d := list.NewDefaultDelegate()
	d.ShowDescription = false
	d.Styles.SelectedTitle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Foreground(lipgloss.Color("#EF6461"))
	d.Styles.NormalTitle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("192")).PaddingLeft(3)

	l := list.New(items, d, 0, 0)
	l.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("#cacccf")).Background(lipgloss.Color("#EF6461"))

	l.SetShowStatusBar(false)

	m := episodeModel{
		episodeList: l,
		err:         nil,
		loading:     false,
	}
	return &m
}

func (m episodeModel) Init() (tea.Cmd) {
	m.loading = true
	return getEpisodeResults(m.animeId)
}

func (m episodeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" { return m, tea.Quit }
		if msg.String() == "enter" {
			m.episodeId = m.episodeList.SelectedItem().(episodes).id
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.episodeList.SetSize(msg.Width-h, msg.Height-v)

	case epRes:
		m.loading = false
		epResults := []episodes(msg)
		
		eList := []list.Item{}
		for _, episode := range epResults {
			eList = append(eList, episodes{
				id:    episode.id,
				epNum: episode.epNum,
			})
		}
		m.episodeList.SetItems(eList)
		return m, nil
	case errMsg:
		m.err = msg
	}

	var cmd tea.Cmd
	m.episodeList, cmd = m.episodeList.Update(msg)
	return m, cmd
	
}

func (m episodeModel) View() string  {
	if m.loading {
		return "Fetching episode Results..."
	}
	return docStyle.Render(m.episodeList.View())
}


func getEpisodeResults(Id string) tea.Cmd {
	return func() tea.Msg {
		episodeList, err := getEpisodeList(Id)
		if err != nil {
			return errMsg{err}
		}
		log.Printf("about to return Results")
		return epRes(episodeList)
	}
}


func getEpisodeList(aniID string) ([]episodes, error) {
	epUrl := fmt.Sprintf("%sanime/gogoanime/info/%s", utils.BaseApiUrl, aniID)
	body, err := utils.FetchJsonData(epUrl)
	
	var epStruct JsonsStrcuts.AnimeInfo
	err = json.Unmarshal(body, &epStruct)
	utils.CheckErr(err)	
	
	episodeList := []episodes{}
	for _, source := range epStruct.Episodes {
		episodeList = append(episodeList,
		episodes {
			id:    source.ID,
			epNum: source.Number,
		})
	}
	return episodeList, nil
}

func (m episodeModel) SetAnimeId(id, name string) (episodeModel, tea.Cmd) {
	m.loading = true
	m.animeId = id
	log.Printf("EPISODE ID: %s", id)
	m.episodeList.Title = fmt.Sprintf("Episodes for %s", name)
	return m, getEpisodeResults(id)
}


