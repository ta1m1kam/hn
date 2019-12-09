package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"strconv"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func generateTreeNodes() []*widgets.TreeNode {
	hns := GetHackerNews()
	var nodes []*widgets.TreeNode
	for _, hn := range hns {
		fmt.Println(hn.Score)
		node := widgets.TreeNode{
			Value: nodeValue(hn.Title),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue("Score: " + strconv.Itoa(hn.Score)),
					Nodes: nil,
				},

				{
					Value: nodeValue("Type: " + hn.Type),
					Nodes: nil,
				},
				{
					Value: nodeValue("Author: " + hn.By),
					Nodes: nil,
				},
				{
					Value: nodeValue("cmd+click →  " + hn.Url),
					Nodes: nil,
				},
				{
					Value: nodeValue("Description"),
					Nodes: []*widgets.TreeNode{
						{
							Value: nodeValue(hn.Description),
							Nodes: nil,
						},
					},
				},
			},
		}
		nodes = append(nodes, &node)
	}

	return nodes
}

func main() {

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to init")
	}
	defer ui.Close()

	nodes := generateTreeNodes()

	t := widgets.NewTree()
	t.Title = "Hacker News ClI"
	t.TextStyle = ui.NewStyle(ui.ColorYellow)
	t.WrapText = false
	t.SetNodes(nodes)
	x, y := ui.TerminalDimensions()
	t.SetRect(0, 0, x, y)
	ui.Render(t)
	Keybindings(t)
}
