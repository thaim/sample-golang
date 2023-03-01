package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/hcl/v2/hclparse"
)

func main() {
	p := hclparse.NewParser()

	file, diags := p.ParseHCLFile("sample.hcl")
	if diags.HasErrors() {
		fmt.Println(diags.Error())
		return
	}
	fmt.Printf("%s", file.Bytes)

	// hclパッケージを用いず予めバッファにファイルを読み込んでおく
	var body []byte
	f, err := os.Open("sample2.hcl")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	body, err = io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ParseHCLに読み込み済のバッファとその仮想ファイル名を与える
	file, diags = p.ParseHCL(body, "sampleX.hcl")
	if diags.HasErrors() {
		fmt.Println(diags.Error())
		return
	}
	fmt.Printf("%s", file.Bytes)

	// Filesではファイル名とコンテンツのmapが返る
	mapping := p.Files()
	for k, v := range mapping {
		fmt.Printf("filename: %s\n", k)
		fmt.Printf("contents: %s\n", v.Bytes)
	}

	// Sourcesを使えばhcl.Fileを経由せずにコンテンツを取得できる
	// これはパースエラー時に利用する(hcl.Fileでは不正なファイルを表現できないので)
	src := p.Sources()
	for k, v := range src {
		fmt.Printf("filename: %s\n", k)
		fmt.Printf("contents: %s\n", v)
	}
}
