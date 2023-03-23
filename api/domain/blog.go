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

func NewBlog() (*Blog, error) {
	const layout = "2006-01-02"
	dayStr := time.Now().Format(layout)

	return &Blog{
		title:   dayStr,
		tags:    []string{"Kindergarten"},
		pubDate: dayStr,
		slug:    dayStr,
	}, nil
}

func (blog *Blog) ToText() string {
	text := ""
	text += "---\n"
	text += fmt.Sprintf("title \"%s\"\n", blog.title)
	text += fmt.Sprintf("tags \"%s\"\n", blog.tags)
	text += fmt.Sprintf("pubDate \"%s\"\n", blog.pubDate)
	text += "---\n\n\n"

	return text
}
