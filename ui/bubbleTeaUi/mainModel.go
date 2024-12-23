package bubbleTeaUi

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	searchAPage state = iota
	animePage
)

type MainModel struct {
	state       state
	searchModel tea.Model
	animeModel  tea.Model
	styles      *Styles
	WindowSize tea.WindowSizeMsg
}

func New() tea.Model {
	return MainModel{
		state:       searchAPage,
		searchModel: *NewSearchModel(),
		animeModel:  *NewResultsModel(),
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

			updatedAnimeModel, searchCmd := m.animeModel.(animeModel).SetSearchTerm(searchValue)

			m.animeModel = updatedAnimeModel

			cmds = append(cmds, searchCmd, tea.WindowSize())
		}
		cmd = newCmd

	case animePage:
		newModel, newCmd := m.animeModel.Update(msg)
		m.animeModel = newModel.(animeModel)
		cmd = newCmd
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
	}

	return "fail to display views (form main model)"
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
