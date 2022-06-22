package controllers

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/utils"
	"github.com/sirupsen/logrus"
)

type controller struct{}

func queryWithFallback(s string, defaultValue string) string {
	if s == "" {
		return defaultValue
	}
	return s
}

func logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithFields(logrus.Fields{
			"path":   c.Request.RequestURI,
			"method": c.Request.Method,
			"status": c.Writer.Status(),
		}).Info(time.Now().Format(time.RFC3339))

	}
}

func selectStrategiesCondition(s string) string {
	if s == `` {
		return `debtLimit`
	}
	for _, c := range utils.STRATEGIES_CONDITIONS {
		if strings.EqualFold(c, s) {
			return c
		}
	}
	return `debtLimit`
}
