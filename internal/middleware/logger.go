package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logging() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now().UTC()
		bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back
		fmt.Println(string(bodyBytes))

		context.Next()

		end := time.Now().UTC()

		zap.S().Infow("[YouHeLogInfo] time_interval: "+end.Sub(start).String(),
			"uri", context.Request.RequestURI,
			"method", context.Request.Method,
			"args", context.Request.PostForm,
			"body", string(bodyBytes),
			"from", context.ClientIP())
	}
}
