package cmd

import (
	"fmt"
	"gin-vue-admin/config"
	log "gin-vue-admin/init/initlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RunWindowsServer(Router *gin.Engine) {
	address := fmt.Sprintf(":%d", config.GinVueAdminconfig.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	log.L.Debug("server run success on ", address)
	fmt.Printf("欢迎使用gin-vue-admin"+
		"默认自动化文档地址:http://127.0.0.1%s/swagger/index.html\n\t"+
		"默认前端文件运行地址:http://127.0.0.1:8080", s.Addr)
	_ = s.ListenAndServe()
}
