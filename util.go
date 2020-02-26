package pipeline

import "golang.org/x/net/html"

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
