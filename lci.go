package main

import (
	"log"
	"os"

	lci "github.com/JackieLi565/lci/cli"
	"github.com/JackieLi565/lci/cmd/pattern"
	"github.com/urfave/cli/v2"
)

func main() {
	app := lciMain()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func lciMain() *cli.App {
	lciInstance := lci.NewLCI()

	app := cli.App{
		Name:    "lci",
		Usage:   "Leetcode tracker via the terminal",
		Version: "0.0.1",
		Commands: []*cli.Command{
			pattern.NewPatternCmd(lciInstance),
		},
	}

	return &app
}
