package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
)

func printHeader(fields []string) {
	var header string
	for _, name := range fields {
		header += fmt.Sprintf("%s\t", name)
	}
	fmt.Println(strings.TrimSpace(header))
}

func main() {
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
		cli.StringFlag{
			Name:  "charset, c",
			Value: "",
			Usage: "Set charset of HTML files.",
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

		charset := ctx.String("charset")

		if !ctx.Bool("no-header") {
			printHeader(fields)
		}

		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if isHtml, _ := regexp.MatchString(".*.html?$", path); isHtml {
				var v values
				v.pickOutFrom(path, charset)
				v.print(fields, root)
			}

			return nil
		})
	}
	app.Run(os.Args)
}
