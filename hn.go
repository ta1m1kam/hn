package main

import (
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hn"
	app.Usage = "This is a tool to see 'Hacker News' made with Go"

	app.Commands = []cli.Command{
		{
			Name:    "number",
			Aliases: []string{"n"},
			Usage:   "options for number of Hacker News acquisitions",
			Action: func(c *cli.Context) error {
				number, _ := strconv.Atoi(c.Args().First())
				hnUi(number)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
