package pipeline

import (
	"testing"

	"github.com/microcosm-cc/bluemonday"
)

func TestSanitizationFilterWithCustomPolicy(t *testing.T) {
	policy := bluemonday.NewPolicy()
	policy.AllowElements("p")
	policy.AllowElements("img")
	policy.AllowAttrs("src").OnElements("img")

	pipeline := NewPipeline([]Filter{
		SanitizationFilter{
			Policy: policy,
		},
	})

	html := `<div><p style="color: red">Hello <span><em>w</em>orld</span></p><img src="foo" width="100" /></div>`
	expected := `<p>Hello world</p><img src="foo"/>`

	assertCall(t, pipeline, expected, html)
}

func TestSanitizationFilter(t *testing.T) {
	pipeline := NewPipeline([]Filter{
		SanitizationFilter{},
	})

	html := `<p style="margin: 0pt;"><img alt="" src="https://helloworld.com/images/a4c7e5612772b2429791790c7e54eeba.jpg" width="100px" style="width: 600px; height: 484px;"></p>
	<p style="margin: 0pt;"><span style="font-family: 宋体; font-size: 10.5pt; mso-spacerun: &quot;yes&quot;; mso-ascii-font-family: Calibri; mso-hansi-font-family: Calibri; mso-bidi-font-family: &quot;Times New Roman&quot;; mso-font-kerning: 1.0000pt;"><font color="#000000" face="宋体">美股研究社</font><font color="#000000">1</font><font color="#000000"><font face="宋体">月</font><font face="Calibri">8</font><font face="宋体">日消息，知名投资机构</font></span></p>
	<blockquote>This is blockquote</blockquote>
	<table><tr><th>Foo</th><tr><td width="100">Bar</td></tr></table>
	<ul><li>First line</li><li><strong>S<b>e</b><i>c</i>ond</strong> line</li></ul>
	<ol><li>First line</li><li>Second line</li></ol>`

	expected := `<p><img alt="" src="https://helloworld.com/images/a4c7e5612772b2429791790c7e54eeba.jpg"/></p>
	<p><span>美股研究社1月8日消息，知名投资机构</span></p>
	<blockquote>This is blockquote</blockquote>
	<table><tbody><tr><th>Foo</th></tr><tr><td width="100">Bar</td></tr></tbody></table>
	<ul><li>First line</li><li><strong>S<b>e</b><i>c</i>ond</strong> line</li></ul>
	<ol><li>First line</li><li>Second line</li></ol>`

	assertCall(t, pipeline, expected, html)
}
