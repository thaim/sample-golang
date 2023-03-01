package main

import(
	"fmt"
	"os"

	"github.com/minamijoyo/hcledit/editor"
)

func main() {
	client := editor.NewClient(&editor.Option{InStream: os.Stdin, OutStream: os.Stdout, ErrStream: os.Stderr})

	// bの値を取得するsink
	sink := editor.NewAttributeGetSink("b")
	client.Derive("sample.hcl", sink)
	fmt.Println()

	// a = "xyz"に書き換えるfilter
	filter := editor.NewAttributeSetFilter("a", "xyz")
	// sample.hclにfilterを適用する。update:falseなのでOutStreamに結果を出力する
	client.Edit("sample.hcl", false, filter)
	fmt.Println()

	// attr1 = "val1"の値を追加する
	filter = editor.NewAttributeAppendFilter("attr1", "\"val1\"", true)
	client.Edit("sample.hcl", false, filter)
}
