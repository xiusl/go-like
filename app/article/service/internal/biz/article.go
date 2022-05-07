package biz

import "time"

type ArticleRepo interface {
}

type Article struct {
	Id        int64
	Title     string
	Url       string
	Content   string
	Images    []string
	UserId    int64
	CreatedAt time.Time
}
