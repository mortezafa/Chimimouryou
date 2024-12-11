package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"	
)
func ShowSearchPage() {
	app:= tview.NewApplication()
	searchMenu(app)
}

func searchMenu(app *tview.Application)  {
	
	titleText := tview.NewTextView().
		SetText(`                                                                                                                        
 ▒▒▒▒▒▒▒░    ▒▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▒▒▒▒▒▒░   ░▒▒▒▒▒▒▒▒▒         ░▒▒▒▒▒░   ▒▒     ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒         ░▒▒▒▒▒▒▒▒▒ 
 ▒▒▒░     █        ▒▒░       ░░░░░▒░     █      ░▒▒▒▒ ████████░▒▒▒▒▒  █ ░  ███  ▒▒▒▒▒▒▒▒░░░▒▒▒▒▒▒▒▒▒ ███████▒░▒▒▒▒▒▒▒▒▒ 
 ▒▒▒░▒████████████ ▒▒░███ ██ ░ ░ ░▒░░███████████░▒▒▒░           ░▒▒▒ ██ ░ █▓ ██ ▒▒▒▒▒▒▒░   ░░░░▒▒▒▒░           ▒▒▒▒▒▒▒▒ 
 ▒▒▒░    █         ▒▒     █    █ ░▒░     █       ▒▒▒  ▓████████  ░▒▒ ██ ▒    ██ ░▒▒▒▒▒▒░ █  ░█ ▒▒▒░  █████████  ▒▒▒▒▒▒▒ 
 ▒▒▒▒▒░  █ █████   ▒  ▓█████████    ▒██████████░  ▒▒ ██        █▒ ▒▒ ██ ▒▒▒▒ ██ ░▒▒▒▒▒▒░ █████ ▒▒▒░██▒       ██ ▒▒▒▒▒▒▒ 
 ▒▒▒▒▒░ ██▒     ██ ▒ ██  █    █████     █      ██  ▒    ░░▒▒░  █▒ ▒▒ ██ ▒▒░  ██ ▒▒▒▒     █    ░▒▒▒░   ░░▒▒▒░  █ ▒▒▒▒▒▒▒ 
 ▒▒▒▒▒░         ██ ▒ ██ ██    █░    ▒▒░ █       ██ ▒▒▒▒▒░     ██  ▒▒        ██  ▒▒▒▒ ███████   ░▒▒▒▒▒▒░░     ██ ▒▒▒▒▒▒▒ 
 ▒▒▒▒▒▒ █████████  ▒  ██░   ███  ▒▒▒▒▒░  ████████  ▒▒▒░  ░████   ░▒▒▒▒▒ ████▓  ▒▒▒▒▒ █████  ██░░▒▒▒▒▒   █████   ▒▒▒▒▒▒▒ 
 ▒▒▒▒▒▒           ▒▒░    ░▒     ▒▒▒▒▒▒▒░          ▒▒▒▒░░██     ░▒▒▒▒▒▒▒       ▒▒▒▒▒▒           ░▒▒▒▒▒ ██▓     ░▒▒▒▒▒▒▒▒ 
 ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░    ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒    ░▒▒▒▒▒▒▒▒▒▒▒▒▒ 
                                                                                                                        `).
		SetWrap(false).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorCadetBlue)
	
	titleFlex := tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(titleText, 0, 1, true)
	
	episodesFlex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Episodes"), 0, 2, false)
	
	episodeSection := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleFlex, 11, 1, true).
		AddItem(episodesFlex, 0, 1, false)
	
	animeBox := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Anime"), 0, 1, false)
	
	flex := tview.NewFlex().
		AddItem(animeBox, 0, 1, false).
		AddItem(episodeSection, 0, 2, false)
	
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
	
}
