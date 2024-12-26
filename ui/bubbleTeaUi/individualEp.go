package bubbleTeaUi

import (
	"Chimimouryou/JsonsStrcuts"
	"Chimimouryou/utils"
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os/exec"
)

type indivEpModel struct{
	loading bool
	epID string
	err error
	mpvRan   bool // new
	
}

type episodeLinks struct {
	link string
	quality string
	isM3u8 bool
}

type (
	linkRes []episodeLinks
	epLinkErrMsg struct{ err error }
)

type playLinkMsg struct {
	url string
}
type mpvFinishedMsg struct {
	err error
}

func NewIndivEpModel() *indivEpModel {
	return &indivEpModel{
		loading: false,
		epID:    "",
		err:     nil,
	}
}

func (m indivEpModel) Init() (tea.Cmd) {
	m.loading =  true
	return GetLinkCmd(m.epID)
}

func (m indivEpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

	case linkRes:
		if !m.mpvRan {
			m.mpvRan = true
			m.loading = false
			links := []episodeLinks(msg)

			var url string
			for _, link := range links {
				if link.quality == "1080p" && link.isM3u8 {
					url = link.link
					break
				} else {
					if link.quality == "720p" && link.isM3u8 {
						url = link.link
						break
					}
				}
			}
			return m, func() tea.Msg {
				return playLinkMsg{url: url}
			}
		}
		return m, nil
	case playLinkMsg:
		url := msg.url
		return m, func() tea.Msg {
			err := runMPV(url)
			return mpvFinishedMsg{err: err}
		}
	case mpvFinishedMsg:
		newM := m
		if msg.err != nil {
			newM.err = msg.err
		}
		newM.mpvRan = false
		return newM, nil

	case epLinkErrMsg:
		newM := m
		newM.err = msg.err
		return newM, nil
		
	}
	var cmd tea.Cmd
	return m, cmd

}

func (m indivEpModel) View() string  {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress 'b' to go back.", m.err)
	}
	if m.loading {
		return "Fetching Video..."
	}
	if m.mpvRan {
		return "Playing Video in MPV..."
	}
	return "Press 'b' to go back. (Not currently playing anything)"
}


func GetLinkCmd(epId string) (tea.Cmd) {
	return func() tea.Msg {
		streamList, err := getEpisodeLinks(epId)
		utils.CheckErr(err)
		
		return linkRes(streamList)
	}
}

func (m indivEpModel) playLink(id string) (indivEpModel, tea.Cmd) {
	m.loading = true
	m.epID = id
	return m, GetLinkCmd(id)
}

func getEpisodeLinks(epID string) ([]episodeLinks, error) {
	url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/watch/%s", epID)
	body, err := utils.FetchJsonData(url)
	
	var links JsonsStrcuts.AnimeStreams	
	err = json.Unmarshal(body, &links)
	utils.CheckErr(err)
	
	streamsList := []episodeLinks{}
	
	for _, source := range links.Sources {
		streamsList =  append(streamsList, episodeLinks{
			link:    source.Url,
			quality: source.Quality,
			isM3u8:  source.IsM3U8,
		})
	}
	return streamsList, nil
}

func runMPV(link string) error {
	mpvCommand := exec.Command("mpv", fmt.Sprintf("%s", link), "--fs")

	err := mpvCommand.Run()
	if err != nil {
		log.Printf("Failed to start MPV: %v", err)
		return err
	}

	return nil
}





