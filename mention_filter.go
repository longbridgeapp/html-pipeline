package pipeline

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var (
	// A-Za-z0-9\p{Han}
	mentionNameFormat = `\w\p{Han}_-`
)

// MentionFilter mention with @ or # or other prefix
type MentionFilter struct {
	// Mention prefix char, default: @
	Prefix string
	// Format func for format matched names to HTML or other
	Format func(name string) string
	// NamesCallback return matched names
	NamesCallback func(names []string)
	mentionRegexp *regexp.Regexp
}

func (f *MentionFilter) initDefault() {
	if f.Prefix == "" {
		f.Prefix = "@"
	}

	if f.mentionRegexp == nil {
		f.mentionRegexp = regexp.MustCompile(`(^|[^` + mentionNameFormat + `])(` + f.Prefix + `)([` + mentionNameFormat + `]+)`)
	}
}

func (f MentionFilter) Call(doc *goquery.Document) (err error) {
	f.initDefault()

	names := []string{}

	rootNode := doc.Find("body")

	TraverseTextNodes(rootNode.Nodes[0], func(node *html.Node) {
		if !strings.Contains(node.Data, f.Prefix) {
			return
		}

		_names := f.ExtractMentionNames(node.Data)
		// Replace text to links html
		for _, name := range _names {
			nameHTML := f.Format(name)
			node.Type = html.RawNode
			node.Data = strings.ReplaceAll(node.Data, f.Prefix+name, nameHTML)
		}
		names = append(names, _names...)
	})

	// TODO: 基于 Logins 查询出这些人，给这些人发送通知
	if f.NamesCallback != nil {
		f.NamesCallback(names)
	}

	return
}

// ExtractMentionNames 从一段纯文本中提取提及的用户名
func (f MentionFilter) ExtractMentionNames(text string) (names []string) {
	f.initDefault()

	matches := f.mentionRegexp.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		names = append(names, match[3])
	}
	return
}
