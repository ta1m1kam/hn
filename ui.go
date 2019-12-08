package main

import (
	"fmt"
	"github.com/rivo/tview"
)

func main() {
	hns := GetHackerNews()

	app := tview.NewApplication()
	list := tview.NewList()
	for i, s := range hns {
		fmt.Printf("%T", i)
		list.AddItem(s.Title, "", nil, nil)
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	list.SetBorder(true)
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
