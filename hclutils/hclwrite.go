package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()

	// foo = "bar" を構築する
	rootBody.SetAttributeValue("foo", cty.StringVal("bar"))

	// 以下のHCLを構築する
	// resource "aws_s3_bucket" "this" {
	//   bucket = "sample-bucket-name"
	//   tags = {
	//     baz = true
	//     foo = 10
	//   }
	// }
	block := hclwrite.NewBlock("resource", []string{"aws_s3_bucket", "this"})
	blockBody := block.Body()
	blockBody.SetAttributeValue("bucket", cty.StringVal("sample-bucket-name"))
	blockBody.SetAttributeValue("tags", cty.ObjectVal(map[string]cty.Value{
		"baz": cty.True,
		"foo": cty.NumberIntVal(10),
	}))
	rootBody.AppendBlock(block)

	fmt.Printf("%s", f.Bytes())
}
