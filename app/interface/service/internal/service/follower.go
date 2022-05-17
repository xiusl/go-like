package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"go-like/pkg/xgin"
	"strconv"
)

func (srv *InterfaceService) followingUser(ctx *gin.Context) (interface{}, string, int) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return xgin.ResultFail("参数错误")
	}

	curUser := xgin.GetCurrentUid(ctx)
	if err := srv.userUc.FollowingUser(ctx, curUser, uid); err != nil {
		return xgin.ResultError(err)
	}

	return xgin.ResultSuccess(true)
}

func (srv *InterfaceService) userFollowers(ctx *gin.Context) (interface{}, string, int) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return xgin.ResultFail("参数错误")
	}
	users, err := srv.userUc.GetFollowers(generateContext(ctx), uid)
	if err != nil {
		return xgin.ResultError(err)
	}
	return xgin.ResultSuccess(packUsers(users, true))
}

func (srv *InterfaceService) userFollowings(ctx *gin.Context) (interface{}, string, int) {
	id := ctx.Param("id")
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return xgin.ResultFail("参数错误")
	}
	users, err := srv.userUc.GetFollowings(generateContext(ctx), uid)
	if err != nil {
		return xgin.ResultError(err)
	}
	return xgin.ResultSuccess(packUsers(users, true))
}

func generateContext(ctx *gin.Context) context.Context {
	log.NewHelper(log.DefaultLogger).Infof("current id %d", xgin.GetCurrentUid(ctx))
	kv := []string{
		"user_id", strconv.FormatInt(xgin.GetCurrentUid(ctx), 10),
	}
	return metadata.AppendToClientContext(ctx, kv...)
}
