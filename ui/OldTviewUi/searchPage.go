package OldTviewUi

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
		AddItem(titleText, 0, 1, false)
	
	
	episodeList := tview.NewList().
		AddItem("test 1", "", '◆', nil).
		AddItem("test 1", "", '◆', nil).
		AddItem("test 1", "", '◆', nil).
		AddItem("test 1", "", '◆', nil).
		SetShortcutColor(tcell.ColorPurple)

	episodeFrame := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false). // Empty space above the list
		AddItem(tview.NewBox().
			SetBorder(true).
			SetBorderColor(tcell.ColorCadetBlue).
			SetTitle("Episodes"), 0, 1, false). // Box for border and title
		AddItem(episodeList, 0, 4, true) // Embed the list inside the box

		
	episodeSection := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleFlex, 11, 1, false).
		AddItem(episodeFrame, 0, 2, true)
	
	animeBox := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorCadetBlue).SetTitle("Anime"), 0, 1, false)
	
	flex := tview.NewFlex().
		AddItem(animeBox, 0, 1, false).
		AddItem(episodeSection, 0, 2, false)
	
	if err := app.SetRoot(flex, true).SetFocus(episodeList).Run(); err != nil {
		panic(err)
	}
	
}
