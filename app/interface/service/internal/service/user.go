package service

import (
	"github.com/gin-gonic/gin"
	"go-like/pkg/xgin"
)

func (srv *InterfaceService) mapUser() {
	user := srv.Group("/user")
	user.Use(xgin.AuthMiddleware(srv))
	{
		xgin.POST(user, "", srv.user)
		xgin.GET(user, "/auth", srv.userAuth)
		xgin.POST(user, "/password", srv.userPassword)
	}
}

func (srv *InterfaceService) user(gin *gin.Context) (interface{}, string, int) {
	return nil, "", 0
}

func (srv *InterfaceService) userPassword(gin *gin.Context) (interface{}, string, int) {
	return nil, "", 0
}

func (srv *InterfaceService) userAuth(gin *gin.Context) (interface{}, string, int) {
	return nil, "", 0
}
