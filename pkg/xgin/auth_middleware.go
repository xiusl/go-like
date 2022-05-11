package xgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-like/pkg/security"
	"strings"
)

type AuthServer interface {
	VerifyToken(context.Context, string) (int64, error)
}

func unAuth(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(200, result(nil, "请登录", 403))
}

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(security.AuthorizationHeaderKey)
		if len(authorizationHeader) > 0 &&
			strings.Contains(authorizationHeader, "Bearer ") {
			token := strings.Replace(authorizationHeader, "Bearer ", "", 1)
			uid, _, _, _ := security.ParseToken(token)
			ctx.Set(security.ContextUserKey, uid)
		}
		ctx.Next()
	}
}

func AuthMiddleware(as AuthServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(security.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			unAuth(ctx)
			return
		}
		if !strings.Contains(authorizationHeader, "Bearer ") {
			unAuth(ctx)
			return
		}
		token := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		uid, err := as.VerifyToken(ctx, token)
		if err != nil || uid == 0 {
			unAuth(ctx)
			return
		}
		ctx.Set(security.ContextUserKey, uid)
		ctx.Next()
	}
}

func GetCurrentUid(ctx *gin.Context) int64 {
	if v, ex := ctx.Get(security.ContextUserKey); ex {
		if id, ok := v.(int64); ok {
			return id
		}
	}
	return 0
}
