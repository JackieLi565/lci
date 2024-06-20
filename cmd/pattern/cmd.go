package pattern

import (
	lci "github.com/JackieLi565/lci/cli"
	"github.com/urfave/cli/v2"
)

func NewPatternCmd(instance lci.LCI) *cli.Command {
	return &cli.Command{
		Name:        "pattern",
		Usage:       "Manage patterns",
		Description: "The 'pattern' command allows you to manage patterns, including adding, removing, or listing them.",

		Subcommands: []*cli.Command{
			newAddPatternCmd(instance),
			newListPatternCmd(instance),
			newRemovePatternCmd(instance),
		},

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Enable verbose mode",
			},
		},

		Action: func(c *cli.Context) error {
			cli.ShowSubcommandHelp(c)
			return nil
		},
	}
}
