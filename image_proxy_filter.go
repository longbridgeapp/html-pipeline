package pipeline

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ImageProxyFilter replace img src for use image proxy
type ImageProxyFilter struct {
	// IgnoreHosts, Host list that will ignore, ["your-host.com", "your-assets.com"]
	IgnoreHosts []string
	// Formatter method with
	Formatter func(src string) string
}

// Call render
func (f ImageProxyFilter) Call(doc *goquery.Document) (err error) {
	doc.Find("img").Each(func(i int, node *goquery.Selection) {
		src := node.AttrOr("src", "")
		if !f.IsIgnoreHost(src) {
			newSrc := f.Formatter(src)
			node.SetAttr("src", newSrc)
		}
	})

	return
}

func (f ImageProxyFilter) IsIgnoreHost(src string) bool {
	if src == "" {
		return false
	}

	src = strings.ToLower(src)
	srcURL, err := url.Parse(src)
	if err != nil {
		return false
	}

	for _, host := range f.IgnoreHosts {
		host = strings.Replace(host, "*.", "", 1)
		if strings.HasSuffix(srcURL.Hostname(), strings.ToLower(host)) {
			return true
		}
	}

	return false
}
