package bubbleTeaUi

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	searchAPage state = iota
	animePage
	episodePage
	link
)

type MainModel struct {
	state       state
	searchModel tea.Model
	animeModel  tea.Model
	episodeModel tea.Model
	linkModel tea.Model
	styles      *Styles
	WindowSize tea.WindowSizeMsg
}

func New() tea.Model {
	return MainModel{
		state:       searchAPage,
		searchModel: *NewSearchModel(),
		animeModel:  *NewResultsModel(),
		episodeModel: *NewEpModel(),
		linkModel: *NewIndivEpModel(),
		styles:      DefaultStyles(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowSize = msg
	}

	switch m.state {
	case searchAPage:
		newModel, newCmd := m.searchModel.Update(msg)
		m.searchModel = newModel.(searchModel)
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
			m.state = animePage
			searchValue := m.searchModel.(searchModel).output

			m.animeModel, cmd = m.animeModel.(animeModel).SetSearchTerm(searchValue)
			cmds = append(cmds, cmd, tea.WindowSize())
		}
		cmd = newCmd

	case animePage:
		newModel, newCmd := m.animeModel.Update(msg)
		m.animeModel = newModel.(animeModel)
		cmd = newCmd

		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
			m.state = episodePage
			animeId := m.animeModel.(animeModel).animeID
			animeName := m.animeModel.(animeModel).selectedAnimeName
			log.Printf("Anime Id from main %s", animeId)
			m.episodeModel, cmd = m.episodeModel.(episodeModel).SetAnimeId(animeId, animeName)
			cmds = append(cmds, cmd, tea.WindowSize())
		}
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "ctrl+b" || keyMsg.String() == "left" {
			m.state = searchAPage
			m.searchModel, cmd = m.searchModel.Update(msg)
		}
	case episodePage:
		m.episodeModel, cmd = m.episodeModel.Update(msg)
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
			m.state = link
			epId := m.episodeModel.(episodeModel).episodeId
			
			m.linkModel, cmd = m.linkModel.(indivEpModel).playLink(epId)
			cmds = append(cmds, cmd, tea.WindowSize())
		}
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "ctrl+b" || keyMsg.String() == "left" {
			m.state = animePage
			m.episodeModel, cmd = m.episodeModel.Update(msg)
		}
	case link:
		m.linkModel, cmd = m.linkModel.Update(msg)
		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "ctrl+b" || keyMsg.String() == "left" {
			m.state = episodePage
			m.linkModel, cmd = m.linkModel.Update(msg)
		}
		
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}


func (m MainModel) View() string {
	switch m.state {
	case searchAPage:
		return m.searchModel.View()
	case animePage:
		return m.animeModel.View()
	case episodePage: return m.episodeModel.View()
	case link: return m.linkModel.View()
	}

	return "fail to display views (from main model)"
}

func Main() {
	m := New()

	f, err1 := tea.LogToFile("debug.log", "debug")
	if err1 != nil {
		log.Fatal("err: %w", err1)
	}
	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
}
