package main

import (
	"github.com/rivo/tview"
)

func main() {
	hns := GetHackerNews()

	app := tview.NewApplication()
	list := tview.NewList()
	for _, s := range hns {
		list.AddItem(s.Title, "", 0, nil)
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	list.SetBorder(true)
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
