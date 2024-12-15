package bubbleTeaUi

import tea "github.com/charmbracelet/bubbletea"

type state int

const (
	searchAPage state = iota
	animePage
)

type MainModel struct {
	state       state
	searchModel tea.Model
	animeModel  tea.Model
	activeModel tea.Model
	styles      *Styles
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func View() string {
	return " "
}

func Main() {

}
