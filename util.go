package pipeline

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// TraverseTextNodes map nested node to find all text node
func TraverseTextNodes(node *html.Node, fn func(*html.Node)) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode || node.Type == html.RawNode {
		fn(node)
	}

	cur := node.FirstChild

	for cur != nil {
		next := cur.NextSibling
		TraverseTextNodes(cur, fn)
		cur = next
	}
}

func isHost(hosts []string, src string) bool {
	if src == "" {
		return false
	}

	src = strings.ToLower(src)
	srcURL, err := url.Parse(src)
	if err != nil {
		return false
	}

	for _, host := range hosts {
		host = strings.Replace(host, "*.", "", 1)
		if strings.HasSuffix(srcURL.Hostname(), strings.ToLower(host)) {
			return true
		}
	}

	return false
}
