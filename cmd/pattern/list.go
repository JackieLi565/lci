package pattern

import (
	lci "github.com/JackieLi565/lci/cli"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func newListPatternCmd(instance lci.LCI) *cli.Command {
	cmd := cli.Command{
		Name:        "list",
		Usage:       "list all patterns",
		Description: "the 'list' command displays all patterns and their respective problem count",

		Action: func(ctx *cli.Context) error {
			return listAction(ctx, instance)
		},
	}

	return &cmd
}

func listAction(c *cli.Context, i lci.LCI) error {
	patterns, err := i.DB.ListPatterns()
	if err != nil {
		return err
	}

	var problemSum int
	t := table.NewWriter()
	t.SetOutputMirror(c.App.Writer)
	t.AppendHeader(table.Row{"Pattern", "# of Problems"})

	for _, pattern := range patterns {
		problemSum += pattern.Count
		t.AppendRow(table.Row{pattern.Name, pattern.Count})
	}

	t.AppendFooter(table.Row{"Total", problemSum})
	t.Render()
	return nil
}
