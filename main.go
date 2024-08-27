package main

import (
	"flag"
	"fmt"

	"github.com/sgsavu/tailwindcssparser"
)

func main() {
	var minify = flag.Bool("minify", false, "Minifies the parsed css.")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Example usage: ./tailwindcss-parser --minify \"w-24\"")
		return
	}
	tailwindTags := flag.Arg(0)

	parsedTailwind, err := tailwindcssparser.GetParsedTailwind(tailwindTags, minify)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedTailwind)
}
