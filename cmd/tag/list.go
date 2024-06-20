package tag

import (
	lci "github.com/JackieLi565/lci/cli"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
)

func newListTagCmd(instance lci.LCI) *cli.Command {
	cmd := cli.Command{
		Name:        "list",
		Usage:       "list a new leetcode tag",
		Description: "the 'list' command allows you to list all tags",

		Action: func(ctx *cli.Context) error {
			return listAction(ctx, instance)
		},
	}

	return &cmd
}

func listAction(c *cli.Context, i lci.LCI) error {
	tags, err := i.DB.ListTags()
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(c.App.Writer)
	t.AppendHeader(table.Row{"", "Tag"})

	for _, tag := range tags {
		t.AppendRow(table.Row{"", tag.Name})
	}

	t.AppendFooter(table.Row{"Total", len(tags)})
	t.Render()
	return nil
}
