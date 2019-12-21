package main

import (
	"errors"
	"fmt"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func generateTreeNodes(n int) ([]*widgets.TreeNode, error) {
	hns, err := GetHackerNews(n)
	if err != nil {
		return nil, err
	}
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

	return nodes, nil
}

func hnUi(n int) error {
	if err := ui.Init(); err != nil {
		return errors.New("failed to init")
	}
	defer ui.Close()

	nodes, err := generateTreeNodes(n)
	if err != nil {
		return err
	}

	t := widgets.NewTree()
	t.Title = "Hacker News ClI"
	t.TextStyle = ui.NewStyle(ui.ColorYellow)
	t.WrapText = false
	t.SetNodes(nodes)
	x, y := ui.TerminalDimensions()
	t.SetRect(0, 0, x, y)
	ui.Render(t)
	Keybindings(t)
	return nil
}
