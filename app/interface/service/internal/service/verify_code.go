package service

import (
	"github.com/gin-gonic/gin"
	"go-like/pkg/xgin"
	"strconv"
)

func (srv *InterfaceService) mapVerifyCode() {
	code := srv.Group("/verifycode")
	{
		xgin.GET(code, "", srv.sendVerifyCode)
	}
}

// 发送验证码
// @method get
// @param key string [手机号/邮箱]
// @param biz_type int [业务编号 0 注册/登录 1 修改密码]
// @return success bool 是否成功
func (srv *InterfaceService) sendVerifyCode(ctx *gin.Context) (interface{}, string, int) {
	key := ctx.Query("key")
	if len(key) == 0 {
		return xgin.ResultFail("手机号/邮箱不能为空")
	}
	bizType, err := IntDefaultQuery(ctx, "biz_type")
	if err != nil {
		srv.log.Errorf("转换int数据失败", err)
		return xgin.ResultBad("请求失败")
	}
	err = srv.userUc.SendVerifyCode(ctx, key, bizType)
	if err != nil {
		return xgin.ResultError(err)
	}
	return xgin.ResultSuccess(true)
}

func IntDefaultQuery(ctx *gin.Context, key string) (int64, error) {
	str := ctx.DefaultQuery("biz_type", "0")
	return strconv.ParseInt(str, 10, 64)
}
