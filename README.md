# htmlinfo
Print HTML information ( title, description, keywords ).  
HTMLファイルのタイトルやディスクリプションをタブ文字区切りで出力します。

## Install

Direct downloads are available through the [releases page](https://github.com/daikiueda/htmlinfo/releases/latest).  
実行ファイルは、[release page](https://github.com/daikiueda/htmlinfo/releases/latest)から直接ダウンロードできます。

If you have Go installed on your computer just run go get.  
Goがインストールされている環境であれば、go getでインストールできます。
```Bash
$ go get github.com/daikiueda/htmlinfo
```

## Usage
```
$ htmlinfo [options]
```

### options
* __--dir__, __-d__  
Set top directory of HTML files.

* __--charset__, __-c__  
Set charset of HTML files.

* __--no-header__  
Hide header.
