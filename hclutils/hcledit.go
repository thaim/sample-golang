package main

import(
	"os"

	"github.com/minamijoyo/hcledit/editor"
)

func main() {
	client := editor.NewClient(&editor.Option{InStream: os.Stdin, OutStream: os.Stdout, ErrStream: os.Stderr})

	// a = "xyz"に書き換えるfilter
	filter := editor.NewAttributeSetFilter("a", "xyz")
	// sample.hclにfilterを適用する。update:falseなのでOutStreamに結果を出力する
	client.Edit("sample.hcl", false, filter)
}
