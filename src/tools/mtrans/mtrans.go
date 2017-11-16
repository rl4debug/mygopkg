package main

import (
	"flag"
	"fmt"

	"github.com/rl4debug/mygopkg/src/pkg/hlc"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	var translator = hlc.MakeTranslator("cookie")

	//i prefer use pointers than values in most case
	var request = &hlc.Request{
		Text: "he is",
	}

	response, _ := translator.Translate(request)
	_ = response

	fmt.Println(response)

	flag.String("", "", "")
	fmt.Print("This tool is buiding up")
}
