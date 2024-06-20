package pattern

import (
	"fmt"

	lci "github.com/JackieLi565/lci/cli"
	"github.com/urfave/cli/v2"
)

func newRemovePatternCmd(instance lci.LCI) *cli.Command {
	cmd := cli.Command{
		Name:        "remove",
		Usage:       "remove a pattern",
		UsageText:   "remove [pattern name]",
		Description: "the 'remove' command removes a pattern and its problems",

		Action: func(ctx *cli.Context) error {
			return removeAction(ctx, instance)
		},
	}

	return &cmd
}

func removeAction(c *cli.Context, i lci.LCI) error {
	pattern := c.Args().First()
	if pattern == "" {
		return fmt.Errorf("must provide a pattern name")
	}

	return i.DB.RemovePattern(pattern)
}
