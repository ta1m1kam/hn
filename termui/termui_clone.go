package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initi")
	}
	defer ui.Close()

	// Paragraph
	p := widgets.NewParagraph()
	p.Title = "Text box"
	p.Text = "PRESS q TO QUIT DEMO"
	p.SetRect(0, 0, 50, 3)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorCyan

	// List
	listDate := []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
	}

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = listDate
	l.SetRect(0, 5, 50, 12)
	l.TextStyle.Fg = ui.ColorYellow

	// Tree
	nodes := []*widgets.TreeNode{
		{
			Value: nodeValue("Key 1"),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Key 1.1"),
					Nodes: []*widgets.TreeNode{
						{
							Value: nodeValue("Key 1.1.1"),
							Nodes: nil,
						},
						{
							Value: nodeValue("Key 1.1.2"),
							Nodes: nil,
						},
					},
				},
				{
					Value: nodeValue("Key 1.2"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Key 2"),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Key 2.1"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Key 2.2"),
					Nodes: nil,
				},
				{
					Value: nodeValue("Key 2.3"),
					Nodes: nil,
				},
			},
		},
		{
			Value: nodeValue("Key 3"),
			Nodes: nil,
		},
	}

	t := widgets.NewTree()
	t.TextStyle = ui.NewStyle(ui.ColorYellow)
	t.WrapText = false
	t.SetNodes(nodes)

	t.SetRect(0, 12, 50, 20)

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Drown>":
				l.ScrollDown()
			case "<Up>":
				l.ScrollUp()
			case "j":
				t.ScrollDown()
			case "k":
				t.ScrollUp()
			case "<Enter>":
				t.ToggleExpand()
			case "E":
				t.ExpandAll()
			}
		}
		ui.Render(p, l, t)
	}
}
