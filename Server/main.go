package main

import (
	"gin-vue-admin/cmd"
	"gin-vue-admin/config"
	"gin-vue-admin/init/initRedis"
	"gin-vue-admin/init/initRouter"
	"gin-vue-admin/init/initlog/log"
	"gin-vue-admin/init/mysql"
	"gin-vue-admin/init/registTable"
	"os"
	//"runtime"
)

var (
	mysqlHost = os.Getenv("MYSQLHOST")
	mysqlPort = os.Getenv("MYSQLPORT")
)

func main() {
	if err := log.NewLogger(); err != nil {
		panic(err)
	}
	// 可以通过环境变量来覆盖配置值
	// 未设定有效的环境变量时，使用配置值
	mysqlConfig := config.GinVueAdminconfig.MysqlAdmin
	if mysqlHost != "" && mysqlPort != "" {
		mysqlConfig.Path = mysqlHost + ":" + mysqlPort
	}
	// 链接初始化数据库
	db := mysql.InitMysql(mysqlConfig) // 链接初始化数据库
	if config.GinVueAdminconfig.System.UseMultipoint {
		// 初始化redis服务
		_ = initRedis.InitRedis()
	}
	// 注册数据库表
	registTable.RegistTable(db)
	// 程序结束前关闭数据库链接
	defer mysql.DEFAULTDB.Close()
	// 注册路由
	Router := initRouter.InitRouter()
	Router.Static("/form-generator", "./static/form-generator")
	cmd.RunWindowsServer(Router)
}
