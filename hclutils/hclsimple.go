package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Config struct {
	LogLevel string `hcl:"log_level"`
}

type ConfigInt struct {
	Threshold int `hcl:"power"`
}

func main() {
	var config Config
	var configInt ConfigInt

	// config.hclファイルから読んでデコードする
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err)
		return
	}
	fmt.Printf("Configuration in config.hcl is %#v\n", config)

	// 変数から読んでデコードする
	var varString = `
    log_level = "info"
    `
	err = hclsimple.Decode("config.hcl", []byte(varString), nil, &config)
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err)
		return
	}
	fmt.Printf("Configuration in varString is %#v\n", config)

	var varInt = `
    power = 123
    `
	err = hclsimple.Decode("config.hcl", []byte(varInt), nil, &configInt)
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err)
		return
	}
	fmt.Printf("Configuration in varInt is %#v\n", configInt)

	// 未定義のデータが定義されてるとパースエラーになる
	var varString2 = `
    log_format = "text"
    log_level = "trace"
    log_file = "sample.cfg"
    `
	err = hclsimple.Decode("config.hcl", []byte(varString2), nil, &config)
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err)
	}
	// errにはなるが読み込めている
	fmt.Printf("Configuration in varString2 is %#v\n", config)

	var varJson = `
    {
        "log_level": "error"
    }
    `
	err = hclsimple.Decode("config.json", []byte(varJson), nil, &config)
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err)
	}
	// jsonファイルも読み込める
	fmt.Printf("Configuration in varJson is %#v\n", config)
}
