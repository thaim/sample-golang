hclutils
==========

HCLファイルの操作ライブラリの動作確認

## hclsimple
[hclsimple](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclsimple)はHCLフィアルをパースしてGoの構造体に読み込むパッケージ。
タグ情報付きの構造体を渡して読み込むので、事前にファイルの構造を定義しておく必要がある。

```bash
$ go run hclsimple.go
Configuration in config.hcl is main.Config{LogLevel:"debug"}
Configuration in varString is main.Config{LogLevel:"info"}
Configuration in varInt is main.ConfigInt{Threshold:123}
Failed to load configuration: config.hcl:2,5-15: Unsupported argument; An argument named "log_format" is not expected here., and 1 other diagnostic(s)
Configuration in varString2 is main.Config{LogLevel:"trace"}
Configuration in varJson is main.Config{LogLevel:"error"}
```


## hclwrite
https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclwrite

```bash
$ go run hclwrite.go
string = "bar"
```


## hclparse
https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclparse

```bash
$ go run hclparse.go
a = "aaa"
b = ["x", "y", "z"]
```

## hcledit
https://github.com/minamijoyo/hcledit

```bash
$ go run hcledit.go
["x", "y", "z"]

a = xyz
b = ["x", "y", "z"]

a = "aaa"
b = ["x", "y", "z"]

attr1 = "val1"

a = "aaa"
```
