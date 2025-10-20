package middlewares

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger") {
			c.Next()
		} else {
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			start := time.Now() // start
			path := c.FullPath()

			rawQuery := c.Request.URL.RawQuery
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()

			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			c.Writer = blw
			c.Next()

			params := gin.LogFormatterParams{}
			params.TimeStamp = time.Now() // stop

			params.Latency = params.TimeStamp.Sub(start)
			params.ClientIP = c.ClientIP()
			params.Method = c.Request.Method
			params.StatusCode = c.Writer.Status()
			params.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			params.BodySize = c.Writer.Size()

			if rawQuery != "" {
				path = path + "?" + rawQuery
			}

			params.Path = path

			keys := map[logging.ExtraKey]interface{}{}
			keys[logging.ClientIp] = params.ClientIP
			keys[logging.Method] = params.Method
			keys[logging.Latency] = params.Latency
			keys[logging.StatusCode] = params.StatusCode
			keys[logging.ErrorMessage] = params.ErrorMessage
			keys[logging.BodySize] = params.BodySize
			keys[logging.RequestBody] = string(bodyBytes)
			keys[logging.ResponseBody] = blw.body.String()
			keys[logging.Path] = params.Path

			logger.Info(logging.RequestResponse, logging.Api, "", keys)

		}
	}
}
