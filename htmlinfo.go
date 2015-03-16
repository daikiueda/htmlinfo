package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/codegangsta/cli"
	"regexp"
)

type values struct {
	Path        string
	Title       string
	Description string
	Keywords    string
}

func (v *values) get(key string) string {
	r := reflect.ValueOf(v)
	return reflect.Indirect(r).FieldByName(key).String()
}

func (v *values) print(fields []string, root string) {
	var recode string
	for _, name := range fields {
		val := v.get(name)
		if name == "Path" {
			val, _ = filepath.Rel(root, val)
			val = "/" + val
		}
		recode += fmt.Sprintf("%s\t", val)
	}
	fmt.Println(strings.TrimSpace(recode))
}

func (v *values) pickOutFrom(path string) {

	file, _ := os.Open(path)
	reader := bufio.NewReader(file)
	doc, _ := goquery.NewDocumentFromReader(reader)
	defer file.Close()

	v.Path = path
	v.Title = doc.Find("title").Text()
	v.Description, _ = doc.Find("meta[name=description]").Attr("content")
	v.Keywords, _ = doc.Find("meta[name=Keywords]").Attr("content")
}

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
			printHeader(fields)
		}

		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}

				if isHtml, _ := regexp.MatchString(".*.html?$", path); isHtml {
					var v values
					v.pickOutFrom(path)
					v.print(fields, root)
				}

				return nil
			})
	}
	app.Run(os.Args)
}
