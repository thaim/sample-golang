package main

import (
	"fmt"

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
}
