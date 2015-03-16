package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	var v values
	v.pickOutFrom("test/data/htdocs_utf8/index.html", "")
	if v.get("Title") != "サンプルサイト （+PR＆SEO向けメッセージ）" {
		t.Fail()
	}
}

func TestPickOutFrom(t *testing.T) {
	var v values
	v.pickOutFrom("test/data/htdocs_utf8/index.html", "")
	if v.Title != "サンプルサイト （+PR＆SEO向けメッセージ）" {
		t.Fail()
	}
}

func ExamplePrint() {
	var v values
	v.pickOutFrom("test/data/htdocs_utf8/index.html", "")

	fields := []string{
		"Path",
		"Title",
		"Description",
		"Keywords",
	}
	root := "test/data/htdocs_utf8/"

	v.print(fields, root)

	// Output:
	// /index.html	サンプルサイト （+PR＆SEO向けメッセージ）	サンプルサイトのトップページ
}
