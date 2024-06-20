package main

import (
	"log"
	"os"

	lci "github.com/JackieLi565/lci/cli"
	"github.com/JackieLi565/lci/cmd/pattern"
	"github.com/JackieLi565/lci/cmd/tag"
	"github.com/urfave/cli/v2"
)

func main() {
	app := lciMain()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func lciMain() *cli.App {
	var lciInstance lci.LCI

	app := cli.App{
		Name:    "lci",
		Usage:   "Leetcode tracker via the terminal",
		Version: "0.0.1",
		Commands: []*cli.Command{
			pattern.NewPatternCmd(lciInstance),
			tag.NewTagCmd(lciInstance),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.Args().First() != "config" {
				configPath := ctx.String("config")
				lciInstance = lci.NewLCI(configPath)
			}
			return nil
		},
	}

	return &app
}
