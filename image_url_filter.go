package pipeline

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// ImageURLFilter will match image src and replace with custom format.
type ImageURLFilter struct {
	// IgnoreHosts, Host list that will ignore, ["your-host.com", "your-assets.com"], if empty will match all.
	IgnoreHosts []string
	// MatchHosts, ["some-host.com", "some-assets.com"], host list that match will do format,
	// If present will ignore IgnoreHosts rules,
	// Otherwice will use IgnoreHosts rules.
	MatchHosts []string
	// Format method with
	Format func(src string) string
}

// Call render
func (f ImageURLFilter) Call(doc *goquery.Document) (err error) {
	doc.Find("img").Each(func(i int, node *goquery.Selection) {
		src := node.AttrOr("src", "")
		if src == "" {
			// skip empty src
			return
		}

		// Fix src that not URL Schema
		srcURL, err := url.Parse(src)
		// ignore invaild src
		if err != nil {
			return
		}

		if srcURL.Host == "" {
			// skip relative url
			return
		}

		if srcURL.Scheme == "" {
			srcURL.Scheme = "https"
			src = srcURL.String()
		}

		var matched = false
		if len(f.MatchHosts) > 0 {
			// If has MatchHosts, match first
			matched = isHost(f.MatchHosts, src)
		} else {
			// ignore IgnoreHosts
			matched = !isHost(f.IgnoreHosts, src)
		}

		if matched {
			newSrc := f.Format(src)
			node.SetAttr("src", newSrc)
		}
	})

	return
}
