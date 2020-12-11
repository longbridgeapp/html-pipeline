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
