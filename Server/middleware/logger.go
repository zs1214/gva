package middleware

import (
	"bytes"
	log "gin-vue-admin/init/initlog"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request time
		start := time.Now()
		// request path
		path := c.Request.URL.Path
		logFlag := true
		if strings.Contains(path, "swagger") {
			logFlag = false
		}
		// request ip
		clientIP := c.ClientIP()
		method := c.Request.Method
		// copy request content
		req, _ := httputil.DumpRequest(c.Request, true)
		if logFlag {
			log.L.Debug(
				"Request:", method, &clientIP, path, string(req))
		}
		// replace writer
		cusWriter := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = cusWriter
		// handle request
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		statuscode := c.Writer.Status()
		if logFlag {
			log.L.Debug(
				"Response:",
				statuscode,
				latency,
				cusWriter.body.String())
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
