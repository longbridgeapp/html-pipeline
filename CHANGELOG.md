## 0.6.0

- Add [ImageURLFilter](https://github.com/huacnlee/html-pipeline/blob/master/image_url_filter.go) match `img` to replace with rules like ([imageproxy](https://github.com/willnorris/imageproxy), Ban URL, Thumb version ...).
- Add [ExternalLinkFilter](https://github.com/huacnlee/html-pipeline/blob/master/external_link_filter.go) to match external links to add `rel="nofollow"`, `target="_blank"`.
- Deprecated **ImageProxyFilter**, please use ImageURLFilter.

## 0.5.0

- Add [ImageProxyFilter](https://github.com/huacnlee/html-pipeline/blob/master/image_proxy_filter.go) for match and replace `img` src.

## 0.4.2

- Ensure to remove HTML Tag for avoid XSS in plain mode.

## 0.4.1

- Fix PlainPipeline to use text output for avoid html in plain.

## 0.4.0

- Add `NewPlainPipeline` for render plain text mode.

## 0.3.2

- Fix output unescape will break html attribute value error.

Before:

```go
in = `<object props="{&quot;url&quot;: &quot;https://example.com/a.jpg&quot;}">We don't like 'escape'</object>`
out = `<object props="{"url": "https://example.com/a.jpg"}">We don't like 'escape'</object>`
```

After

```go
in = `<object props="{&quot;url&quot;: &quot;https://example.com/a.jpg&quot;}">We don't like 'escape'</object>`
out = `<object props="{&#34;url&#34;: &#34;https://example.com/a.jpg&#34;}">We don't like 'escape'</object>`
```

## 0.3.1

- Fix filter result without escape string.

  0.3.0

---

- Add AutoCorrectFilter;

  0.2.2

---

- Fix SimpleFormatFilter to remove `\n`;

  0.2.1

---

- Fix MentionFilter for twice use;

  0.2.0

---

- Add MentionFilter;
- Add HTMLEscapeFilter;

  0.1.0

---

- First release.
