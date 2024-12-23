package bubbleTeaUi

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type episodeModel struct{
	episodeList list.Model
	err       error
	loading   bool
}

type episodes struct{
	id string
	epNum int
}

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
	return nil	
}

func (m episodeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" { return m, tea.Quit }
	
	}
	
	return m, nil
}

func (m episodeModel) View() string  {
	return "EPISODE VIEW"
}


