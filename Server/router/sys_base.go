package router

import (
	"gin-vue-admin/controller/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (r gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("regist", api.Regist)
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("captcha", api.Captcha)
		BaseRouter.POST("captcha/:captchaId", api.CaptchaImg)
	}
	return BaseRouter
}
