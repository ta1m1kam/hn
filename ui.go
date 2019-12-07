package main

import (
	"github.com/marcusolsson/tui-go"
	"log"
)

//type mail struct {
//	from string
//}
//
//var mails = []mail{
//	{
//		from: "John Doe <john@doe.com>",
//	},
//	{
//		from: "Jane Doe <jane@doe.com>",
//	},
//}

func main() {
	hns := GetHackerNews()
	inbox := tui.NewTable(0, 0)
	inbox.SetColumnStretch(0, 3)
	inbox.SetColumnStretch(1, 2)
	inbox.SetFocused(true)

	for _, h := range hns {
		inbox.AppendRow(
			tui.NewLabel(h.Title),
		)
	}

	var (
		from = tui.NewLabel("")
	)

	info := tui.NewGrid(0, 0)
	info.AppendRow(tui.NewLabel("From:"), from)

	mail := tui.NewVBox(info)
	mail.SetSizePolicy(tui.Preferred, tui.Expanding)

	inbox.OnSelectionChanged(func(t *tui.Table) {
		m := hns[t.Selected()]
		from.SetText(m.Title)
	})

	inbox.Select(0)

	root := tui.NewVBox(inbox, tui.NewLabel(""), mail)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Shift+Alt+Up", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
