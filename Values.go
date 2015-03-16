package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Values struct {
	Path        string
	Title       string
	Description string
	Keywords    string
}

func (v *Values) get(key string) string {
	r := reflect.ValueOf(v)
	return reflect.Indirect(r).FieldByName(key).String()
}

func (v *Values) print(fields []string, root string) {
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

func (v *Values) pickOutFrom(path string) {

	file, _ := os.Open(path)
	reader := bufio.NewReader(file)
	doc, _ := goquery.NewDocumentFromReader(reader)
	defer file.Close()

	v.Path = path
	v.Title = doc.Find("title").Text()
	v.Description, _ = doc.Find("meta[name=description]").Attr("content")
	v.Keywords, _ = doc.Find("meta[name=Keywords]").Attr("content")
}
