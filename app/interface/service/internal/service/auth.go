package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-like/app/interface/service/internal/biz"
	"go-like/pkg/xgin"
)

func (srv *InterfaceService) mapAuth() {
	auth := srv.Group("/auth")
	{
		xgin.POST(auth, "", srv.auth)
	}
}

type authParams struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

func (r *authParams) valid() (msg string, ok bool) {
	if len(r.Mobile) == 0 {
		msg = "手机号不能为空"
		return
	}
	if len(r.Code) == 0 {
		msg = "验证码不能为空"
		return
	}
	return "", true
}

func parseAuthParams(ctx *gin.Context) (*authParams, error) {
	var p authParams
	if err := ctx.ShouldBindJSON(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func packUser(u *biz.User, simple bool, token string) gin.H {
	r := gin.H{}
	if u == nil {
		return r
	}
	r["id"] = u.Id
	r["name"] = u.Name
	if !simple {
		r["mobile"] = u.Mobile
	}
	if len(token) > 0 {
		r["token"] = token
	}
	return r
}

func (srv *InterfaceService) auth(ctx *gin.Context) (interface{}, string, int) {
	p, err := parseAuthParams(ctx)
	if err != nil {
		return xgin.ResultError(err)
	}
	u, t, err := srv.authUc.Auth(ctx, p.Mobile, p.Code)
	if err != nil {
		return xgin.ResultError(err)
	}
	return xgin.ResultSuccess(packUser(u, false, t))
}

func (srv *InterfaceService) VerifyToken(ctx context.Context, token string) (int64, error) {
	srv.log.Debugf("verify token")
	return 0, nil
}
