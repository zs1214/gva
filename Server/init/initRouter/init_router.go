package initRouter

import (
	_ "gin-vue-admin/docs"
	log "gin-vue-admin/init/initlog"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	var Router = gin.Default()

	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	// 如果不需要日志 请关闭这里
	Router.Use(middleware.Logger())
	log.L.Debug("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	log.L.Debug("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.L.Debug("register swagger handler")
	// 统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	router.InitMenuRouter(ApiGroup)
	router.InitAuthorityRouter(ApiGroup)
	router.InitApiRouter(ApiGroup)
	router.InitFileUploadAndDownloadRouter(ApiGroup)
	router.InitWorkflowRouter(ApiGroup)
	router.InitCasbinRouter(ApiGroup)
	router.InitJwtRouter(ApiGroup)
	router.InitSystemRouter(ApiGroup)
	router.InitCustomerRouter(ApiGroup)
	router.InitAutoCodeRouter(ApiGroup)
	log.L.Info("router register success")
	return Router
}
