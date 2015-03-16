package main

func ExamplePrintAll() {

	root := "test/data/htdocs_sjis"
	fields := []string{
		"Path",
		"Title",
		"Description",
		"Keywords",
	}
	charset := "shiftjis"

	printAll( root, fields, charset )

	// Output:
	// /index.html	サンプルサイト （+PR＆SEO向けメッセージ）	サンプルサイトのトップページ
	// /sample_dir_1/index.html	サンプルカテゴリー1 | サンプルサイト	サンプルサイトのカテゴリー1
	// /sample_dir_1/sample.html	サンプルページ | サンプルカテゴリー1 | サンプルサイト	サンプルサイトのカテゴリー1のサンプルページ
	// /sample_dir_1/sample_subdir/index.html	サンプルサブカテゴリー | サンプルカテゴリー1 | サンプルサイト	サンプルサイトのカテゴリー1のサブカテゴリー
	// /sample_dir_1/sample_subdir/sample_1.html	サンプルページ1 | サンプルサブカテゴリー | サンプルカテゴリー1 | サンプルサイト	サンプルサイトのカテゴリー1のサブカテゴリーのサンプルページ1
	// /sample_dir_1/sample_subdir/sample_2.html	サンプルページ2 | サンプルサブカテゴリー | サンプルカテゴリー1 | サンプルサイト	サンプルサイトのカテゴリー1のサブカテゴリーのサンプルページ2
}

func ExamplePrintHeader() {
	fields := []string{
		"Path",
		"Title",
		"Description",
		"Keywords",
	}
	printHeader(fields)

	// Output:
	// Path	Title	Description	Keywords
}
