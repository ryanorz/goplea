package ch5

import "golang.org/x/net/html"

// ElementCount ...
func ElementCount(n *html.Node, m map[string]int) map[string]int {
	if n == nil {
		return m
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	m = ElementCount(n.FirstChild, m)
	m = ElementCount(n.NextSibling, m)
	return m
}
