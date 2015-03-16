package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"code.google.com/p/mahonia"
	"github.com/PuerkitoBio/goquery"
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

func (v *values) pickOutFrom(path string, charset string) {

	file, _ := os.Open(path)
	reader := bufio.NewReader(file)
	doc, _ := goquery.NewDocumentFromReader(reader)
	defer file.Close()

	v.Path = path
	v.Title = doc.Find("title").Text()
	v.Description, _ = doc.Find("meta[name=description]").Attr("content")
	v.Keywords, _ = doc.Find("meta[name=Keywords]").Attr("content")

	if charset != "" {
		v.Title = decodeText(v.Title, charset)
		v.Description = decodeText(v.Description, charset)
		v.Keywords = decodeText(v.Keywords, charset)
	}
}

func decodeText(text string, charset string) string {
	if charset == "" {
		return text
	}
	return mahonia.NewDecoder(charset).ConvertString(text)
}
