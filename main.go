package main

import (
	"fmt"
	"github.com/fixmenu/go-html-link-parser/link"
	"strings"
)

func main() {
	h := link.GetHtmlTemplate(5)
	r := strings.NewReader(h)
	link, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", link)
}
