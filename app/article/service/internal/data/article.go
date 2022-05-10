package data

import (
	"context"
	"database/sql"
	"errors"
	"go-like/app/article/service/internal/biz"
	"strings"
)

var articleTableSql = ""

var (
	insertArticleSql = "insert into article(id, title, content, url, images, user_id) " +
		"values (?, ?, ?, ?, ?, ?)"
	selectArticleSql = "select id, title, content, url, images, user_id from article where id=? and del=0"
	listArticleSql   = "select id, title, content, url, images, user_id from article where id > ? order by id limit ?"
)

type articleRepo struct {
	data *Data
}

// NewArticleRepo is .
func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		data: data,
	}
}

// Insert is .
func (r *articleRepo) Insert(ctx context.Context, art *biz.Article) (int64, error) {
	images := strings.Join(art.Images, ",")
	art.Id = r.data.GenerateID()
	res, err := r.data.db.ExecContext(ctx, insertArticleSql, art.Id, art.Title,
		art.Content, art.Url, images, art.UserId)
	if err != nil {
		return 0, err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if row == 0 {
		return 0, errors.New("插入文章失败")
	}
	return art.Id, nil
}

// Select is .
func (r *articleRepo) Select(ctx context.Context, id int64) (*biz.Article, error) {
	row := r.data.db.QueryRowContext(ctx, selectArticleSql, id)
	return scanArticle(row)
}

func scanArticle(row *sql.Row) (*biz.Article, error) {
	var a biz.Article
	err := row.Scan(
		&a.Id,
		&a.Title,
		&a.Content,
		&a.Url,
		&a.Images,
		&a.UserId,
	)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
