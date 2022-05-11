package service

import (
	"github.com/gin-gonic/gin"
	"go-like/app/interface/service/internal/biz"
	"go-like/pkg/xgin"
	"strconv"
)

func (srv *InterfaceService) mapArticle() {
	art := srv.Group("/articles")
	{
		xgin.GET(art, "/:id", srv.getArticle)
	}
	art.Use(xgin.AuthMiddleware(srv))
	{
		xgin.POST(art, "", srv.createArticle)

	}
}

type createArticleParam struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func (r *createArticleParam) valid() (msg string, ok bool) {
	if len(r.Title) == 0 {
		msg = "标题不能为空"
		return
	}
	if len(r.Url) == 0 {
		msg = "链接不能为空"
		return
	}
	return "", true
}

func parseCreateArticleParam(ctx *gin.Context) (*createArticleParam, error) {
	var p createArticleParam
	if err := ctx.ShouldBindJSON(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func packArticle(art *biz.Article, u *biz.User) gin.H {
	data := gin.H{}
	if art == nil || u == nil {
		return data
	}
	data["id"] = art.Id
	data["url"] = art.Url
	data["title"] = art.Title
	data["user"] = packUser(u, true, "")
	return data
}

func (srv *InterfaceService) createArticle(ctx *gin.Context) (interface{}, string, int) {
	p, err := parseCreateArticleParam(ctx)
	if err != nil {
		return xgin.ResultError(err)
	}
	if msg, ok := p.valid(); !ok {
		return xgin.ResultFail(msg)
	}

	uid := xgin.GetCurrentUid(ctx)
	art, u, err := srv.artUc.CreateArticle(ctx, uid, p.Title, p.Url)
	if err != nil {
		return xgin.ResultError(err)
	}
	return xgin.ResultSuccess(packArticle(art, u))
}

func (srv *InterfaceService) getArticle(ctx *gin.Context) (interface{}, string, int) {
	id := ctx.Param("id")
	if len(id) == 0 {
		return xgin.ResultFail("id is null")
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return xgin.ResultError(err)
	}

	uid := xgin.GetCurrentUid(ctx)
	art, u, err := srv.artUc.GetArticle(ctx, uid, aid)
	if err != nil {
		return xgin.ResultError(err)
	}
	return xgin.ResultSuccess(packArticle(art, u))
}
