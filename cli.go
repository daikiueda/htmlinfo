package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) printHeader(fields []string) {
	var header string
	for _, name := range fields {
		header += fmt.Sprintf("%s\t", name)
	}
	fmt.Println(strings.TrimSpace(header))
}

func (c *CLI) Run(args []string) {
	app := cli.NewApp()
	app.Name = "htmlinfo"
	app.Version = "0.0.1"
	app.Usage = "Print HTML info ( title, description, keywords )"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir, d",
			Value: ".",
			Usage: "Set top directory of HTML files.",
		},
		cli.BoolFlag{
			Name:  "no-header",
			Usage: "Hide header.",
		},
	}
	app.Action = func(ctx *cli.Context) {

		root := ctx.String("dir")

		fields := []string{
			"Path",
			"Title",
			"Description",
			"Keywords",
		}

		if !ctx.Bool("no-header") {
			c.printHeader(fields)
		}

		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if isHtml, _ := regexp.MatchString(".*.html?$", path); isHtml {
				var v Values
				v.pickOutFrom(path)
				v.print(fields, root)
			}

			return nil
		})
	}
	app.Run(args)
}
