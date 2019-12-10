package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hn"
	app.Usage = "This is a tool to see 'Hacker News' made with Go"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "number, n",
			Value: 10,
			Usage: "option for number of Hacker News acquisitions",
		},
	}

	app.Action = func(c *cli.Context) error {
		hnUi(c.Int("number"))

		return nil
	}
	app.Run(os.Args)
}
