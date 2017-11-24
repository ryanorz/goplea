package ch5

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// PrintText ...
func PrintText(n *html.Node) {
	if n == nil || n.Data == "script" || n.Data == "style" {
		return
	}
	if n.Type == html.TextNode {
		text := strings.Trim(n.Data, " \t\n")
		if len(text) > 0 {
			fmt.Println(n.Data)
		}
	}
	PrintText(n.FirstChild)
	PrintText(n.NextSibling)
}
