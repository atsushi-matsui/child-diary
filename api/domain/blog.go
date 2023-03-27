package domain

import (
	"fmt"
	"time"
)

type Blog struct {
	title   string
	tags    []string
	pubDate string
	slug    string
}

func NewBlog(title string, slug string) (*Blog, error) {
	const layout = "2006-01-02"
	dayStr := time.Now().Format(layout)

	return &Blog{
		title:   title,
		tags:    []string{"Kindergarten"},
		pubDate: dayStr,
		slug:    slug,
	}, nil
}

func (blog *Blog) ToText() string {
	text := ""
	text += "---\n"
	text += fmt.Sprintf("title: \"%s\"\n", blog.title)
	text += fmt.Sprintf("tags: %s\n", blog.tags)
	text += fmt.Sprintf("pubDate: %s\n", blog.pubDate)
	text += fmt.Sprintf("slug: \"%s\"\n", blog.slug)
	text += "---\n\n\n"

	return text
}
