package tag

import (
	lci "github.com/JackieLi565/lci/cli"
	"github.com/urfave/cli/v2"
)

func NewTagCmd(instance lci.LCI) *cli.Command {
	return &cli.Command{
		Name:        "tag",
		Usage:       "Manage tags",
		Description: "The 'tag' command allows you to manage patterns, including adding, removing, or listing them.",

		Subcommands: []*cli.Command{
			newAddTagCmd(instance),
			newListTagCmd(instance),
			newRemoveTagCmd(instance),
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
