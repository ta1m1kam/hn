package main

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
)

const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Modal
	app  *tview.Application
)

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

func updateTime() {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
			view.SetText(currentTimeString())
		})
	}
}

func main() {
	app = tview.NewApplication()
	view = tview.NewModal().
		SetText(currentTimeString()).
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

	go updateTime()
	if err := app.SetRoot(view, false).Run(); err != nil {
		panic(err)
	}
}
