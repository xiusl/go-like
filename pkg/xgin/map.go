package xgin

import "github.com/gin-gonic/gin"

func GET(r *gin.RouterGroup, path string, handle ginHandlerFunc) {
	r.GET(path, BaseResponse(handle))
}

func POST(r *gin.RouterGroup, path string, handle ginHandlerFunc) {
	r.POST(path, BaseResponse(handle))
}

func PUT(r *gin.RouterGroup, path string, handle ginHandlerFunc) {
	r.PUT(path, BaseResponse(handle))
}

func DELETE(r *gin.RouterGroup, path string, handle ginHandlerFunc) {
	r.DELETE(path, BaseResponse(handle))
}

func PATCH(r *gin.RouterGroup, path string, handle ginHandlerFunc) {
	r.PATCH(path, BaseResponse(handle))
}
