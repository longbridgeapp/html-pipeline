package pipeline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMentionFilterWithPlainText(t *testing.T) {
	text := `This is a @test_huacn-lee of some cool @中文名称 features that @mi_asd be
	@use-ful but @don't. look at this email@address.com. @bla! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the ### comment blocks.
	But #msft is cool.`

	expected := `This is a <mention>test_huacn-lee</mention> of some cool <mention>中文名称</mention> features that <mention>mi_asd</mention> be
	<mention>use-ful</mention> but <mention>don</mention>'t. look at this email@address.com. <mention>bla</mention>! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the ### comment blocks.
	But #msft is cool.`

	pipe := NewPipeline([]Filter{
		MentionFilter{
			Format: func(name string) string {
				return fmt.Sprintf(`<mention>%s</mention>`, name)
			},
		},
	})

	out, err := pipe.Call(text)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestMentionFilterWithHTML(t *testing.T) {
	text := `<p>This is a @test_huacn-lee of some cool @中文名称 features.</p>
	<p>that <span>@mi_asd</span> <strong>@use-ful but @don't. look</strong> at this email@address.com. </p>
	<p>@bla! I like #nylas but I don't like to go to this apple.com?a#url.</p>
	<p>I also don't like the ### comment blocks. But #msft is cool.</p>`

	expected := `<p>This is a <a href="https://twitter.com/test_huacn-lee">@test_huacn-lee</a> of some cool <a href="https://twitter.com/中文名称">@中文名称</a> features.</p>
	<p>that <span><a href="https://twitter.com/mi_asd">@mi_asd</a></span> <strong><a href="https://twitter.com/use-ful">@use-ful</a> but <a href="https://twitter.com/don">@don</a>'t. look</strong> at this email@address.com. </p>
	<p><a href="https://twitter.com/bla">@bla</a>! I like #nylas but I don't like to go to this apple.com?a#url.</p>
	<p>I also don&#39;t like the ### comment blocks. But #msft is cool.</p>`

	matchedNames := []string{}
	mentionfilter := MentionFilter{
		Format: func(name string) string {
			return fmt.Sprintf(`<a href="https://twitter.com/%s">@%s</a>`, name, name)
		},
		NamesCallback: func(names []string) {
			matchedNames = names
		},
	}

	pipe := NewPipeline([]Filter{
		mentionfilter,
	})

	out, err := pipe.Call(text)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
	assert.Equal(t, mentionfilter.ExtractMentionNames(text), matchedNames)
}

func TestMentionFilterTwice(t *testing.T) {
	cases := [][]string{
		[]string{
			`This is #html-pipeline example, created by @huacnlee at 2020.`,
			`This is <hashtag>#html-pipeline</hashtag> example, created by <mention>@huacnlee</mention> at 2020.`,
		},
		[]string{
			`<p>This is <em>#html-pipeline</em> example, created by <strong>@huacnlee</strong> at 2020.</p>`,
			`<p>This is <em><hashtag>#html-pipeline</hashtag></em> example, created by <strong><mention>@huacnlee</mention></strong> at 2020.</p>`,
		},
	}

	hashTags := []string{}
	mentionNames := []string{}

	pipe := NewPipeline([]Filter{
		MentionFilter{
			Prefix: "#",
			Format: func(name string) string {
				return "<hashtag>#" + name + "</hashtag>"
			},
			NamesCallback: func(names []string) {
				hashTags = names
			},
		},
		MentionFilter{
			Format: func(name string) string {
				return "<mention>@" + name + "</mention>"
			},
			NamesCallback: func(names []string) {
				mentionNames = names
			},
		},
	})

	for _, item := range cases {
		text := item[0]
		expected := item[1]

		out, err := pipe.Call(text)
		assert.NoError(t, err)
		assert.Equal(t, []string{"html-pipeline"}, hashTags)
		assert.Equal(t, []string{"huacnlee"}, mentionNames)
		assert.Equal(t, expected, out)
	}

}

func TestExtractMentionNames(t *testing.T) {
	text := `@huacnlee This is a @test_huacn-lee of some cool @中文名称 features that @mi_asd be
	@use-ful but @don't. look at this email@address.com. @bla! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the ### comment blocks.
	But #msft is cool.`

	expectedNames := []string{"huacnlee", "test_huacn-lee", "中文名称", "mi_asd", "use-ful", "don", "bla"}

	mentionFilter := MentionFilter{}

	names := mentionFilter.ExtractMentionNames(text)
	assert.Equal(t, expectedNames, names)
}

func TestExtractMentionNamesWithHashTag(t *testing.T) {
	text := `#huacnlee This is a #test_huacn-lee of some cool #中文名称 features that #mi_asd be
	#use-ful but #don't. look at this email#address.com. #bla! I like #nylas but I don't
	like to go to this apple.com?a#url. I also don't like the comment blocks.
	But #msft is cool.`

	expectedNames := []string{"huacnlee", "test_huacn-lee", "中文名称", "mi_asd", "use-ful", "don", "bla",
		"nylas", "msft"}

	mentionFilter := MentionFilter{Prefix: "#"}

	names := mentionFilter.ExtractMentionNames(text)
	assert.Equal(t, expectedNames, names)
}
