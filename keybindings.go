package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Keybindings do main loop.
func Keybindings(t *widgets.Tree) {
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "k":
			t.ScrollUp()
		case "j":
			t.ScrollDown()
		case "E":
			t.ExpandAll()
		case "C":
			t.CollapseAll()
		case "<Enter>":
			t.ToggleExpand()
		}
		ui.Render(t)
	}
}
