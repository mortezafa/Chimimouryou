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
	searchModel searchModel
	animeModel  animeModel
	activeModel tea.Model
	styles      *Styles
}

func New() tea.Model {
	return MainModel{
		state:       searchAPage,
		searchModel: *NewSearchModel(),
		animeModel:  *NewResultsModel(),
		activeModel: NewSearchModel(),
		styles:      DefaultStyles(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.activeModel.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("insdie Main")
	switch m.state {
	case searchAPage:
		log.Printf("Insde SearchAPage")
		model, cmd := m.searchModel.Update(msg)
		m.searchModel = model.(searchModel)

		if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.String() == "enter" {
			m.state = animePage
			m.activeModel = m.animeModel
			return m, m.activeModel.Init()
		}
		return m, cmd
	}
	return m, nil
}

func (m MainModel) View() string {
	return m.activeModel.View()
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
