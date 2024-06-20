package pattern

import (
	"fmt"

	lci "github.com/JackieLi565/lci/cli"
	"github.com/urfave/cli/v2"
)

func newAddPatternCmd(instance lci.LCI) *cli.Command {
	cmd := cli.Command{
		Name:        "add",
		Usage:       "add a new leetcode pattern",
		Description: "the 'add' command allows you to add a new pattern partition",

		Action: func(ctx *cli.Context) error {
			return addAction(ctx, instance)
		},
	}

	return &cmd
}

func addAction(c *cli.Context, i lci.LCI) error {
	name := c.Args().First()

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Adding a new %q pattern\n", name)
	}

	if name == "" {
		return fmt.Errorf("please provide a valid pattern name")
	}

	err := i.DB.AddPattern(name)
	return err
}
