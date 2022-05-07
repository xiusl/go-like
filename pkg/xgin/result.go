package xgin

import "github.com/gin-gonic/gin"

type ginHandlerFunc func(ctx *gin.Context) (interface{}, string, int)

func BaseResponse(handle ginHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, msg, code := handle(ctx)
		ctx.JSON(200, result(data, msg, code))
	}
}

func result(data interface{}, msg string, code int) gin.H {
	return gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}

func ResultBad(message string) (interface{}, string, int) {
	return nil, message, 400
}

func ResultFail(message string) (interface{}, string, int) {
	return nil, message, 500
}

func ResultError(err error) (interface{}, string, int) {
	return ResultBad(err.Error())
}

func ResultSuccess(data interface{}) (interface{}, string, int) {
	return data, "", 200
}
