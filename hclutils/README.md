hclutils
==========

HCLファイルの操作ライブラリの動作確認

## hclsimple
[hclsimple](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclsimple)はHCLフィアルをパースしてGoの構造体に読み込むパッケージ。
タグ情報付きの構造体を渡して読み込むので、事前にファイルの構造を定義しておく必要がある。
また、[jsonパッケージ](https://pkg.go.dev/github.com/hashicorp/hcl/v2/json)を用いてjsonファイルを読み込むこともできる

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
[hclwrite](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclwrite)はHCLデータを構築するために利用するパッケージ。
`key = value` 形式のAttribute、`res {...}` 形式のBlock、複数のAttributeやBlockで構成されるBodyでHCLデータを構築する。

```bash
$ go run hclwrite.go
foo = "bar"
resource "aws_s3_bucket" "this" {
  bucket = "sample-bucket-name"
  tags = {
    baz = true
    foo = 10
  }
}
```


## hclparse
[hclparse](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclparse)はファイルをパースしてHCLのシンタックスチェックを行う。
また、パースしたファイルとその内容の対応関係を管理する。
シンタックスチェック自体はhclsyntaxで行うが、基本的にはhclparse経由で利用する。

```bash
$ go run hclparse.go
a = "aaa"
b = ["x", "y", "z"]
```


## gohcl
[gohcl](https://pkg.go.dev/github.com/hashicorp/hcl/v2/gohcl) はHCLデータをパースしてGoの構造体に読み込むパッケージ。
低レベルAPIで、hclsimpleなどはgohclを呼び出すことでパースを実現している。
hclsimpleはファイルや[]byteのデータを読み込んでパースするが、
gohclはhcl.Bodyからデータをパースしたり、hclwrite.Bodyにパースしたデータを反映する。


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
