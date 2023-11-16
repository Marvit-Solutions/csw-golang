package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		entry := logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"uri":    c.Request.RequestURI,
			"status": c.Writer.Status(),
		})

		entry.Info("Request processed", "latency", latency)
	}
}
