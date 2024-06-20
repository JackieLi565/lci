package tag

import (
	"fmt"

	lci "github.com/JackieLi565/lci/cli"
	"github.com/urfave/cli/v2"
)

func newRemoveTagCmd(instance lci.LCI) *cli.Command {
	cmd := cli.Command{
		Name:        "remove",
		Usage:       "remove a tag",
		UsageText:   "remove [tag name]",
		Description: "the 'remove' command removes a tag",

		Action: func(ctx *cli.Context) error {
			return removeAction(ctx, instance)
		},
	}

	return &cmd
}

func removeAction(c *cli.Context, i lci.LCI) error {
	tag := c.Args().First()
	if tag == "" {
		return fmt.Errorf("must provide a tag name")
	}

	return i.DB.RemoveTag(tag)
}
