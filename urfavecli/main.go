package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", "))
}

func (e EnumValue) String() string {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}

func main() {
	app := &cli.App{
		Action: sample,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "selected language [english, japanese]",
			},
			&cli.GenericFlag{
				Name: "format, f",
				Value: &EnumValue{
					Enum:    []string{"json", "plist", "xml"},
					Default: "json",
				},
				// Usage: "output in json, plist or xml format",
			},
			&cli.StringFlag{
				Name:  "protocol",
				Value: "https",
				Action: func(ctx *cli.Context, v string) error {
					valid := []string{"http", "https"}
					if !slices.Contains(valid, v) {
						return fmt.Errorf("invalid protocol: %v (must be one of %v)", v, valid)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sample(ctx *cli.Context) error {
	lang := ctx.String("lang")
	switch lang {
	case "english", "japanese":
		fmt.Printf("lang: %s\n", lang)
	default:
		cli.ShowAppHelp(ctx)
		return fmt.Errorf("invalid lang specified: %s\n", lang)
	}

	format := ctx.String("format")
	fmt.Printf("format: %s\n", format)

	protocol := ctx.String("protocol")
	fmt.Printf("protocol: %s\n", protocol)

	return nil
}
