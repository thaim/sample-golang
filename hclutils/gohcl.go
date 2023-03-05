package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
	LogLevel string `hcl:"log_level"`
}

func main() {
	// DecodeBody
	fmt.Println("# DecodeBody")
	sampleDecodeBody()

	// DecodeExpressionの例

	// EncodeIntoBody
	fmt.Println("\n\n# EncodeIntoBody")
	sampleEncodeIntoBody()

	// EncodeAsBlock

	// ImpliedBodySchema
}

// hcl.Bodyを読み込み構造体にデータを読み込む
func sampleDecodeBody() {
	var config Config
	var varConfig = `
    log_level = "info"
    `

	file, diags := hclsyntax.ParseConfig([]byte(varConfig), "config.hcl", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Printf("failed to parse file %s\n", diags)
		return
	}
	diags = gohcl.DecodeBody(file.Body, nil, &config)
	if diags.HasErrors() {
		fmt.Printf("failed to decode %s\n", diags)
		return
	}
	fmt.Printf("config file is %#v", config)
}


// EncodeIntoBodyは構造体のデータを hclwrite.Bodyに読み込む
func sampleEncodeIntoBody() {
	config := Config{
		LogLevel: "trace",
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&config, f.Body())

	fmt.Printf("%s", f.Bytes())
}
