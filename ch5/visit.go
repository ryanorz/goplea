package ch5

import "golang.org/x/net/html"

// visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = Visit(links, n.FirstChild)
	links = Visit(links, n.NextSibling)
	return links
}
