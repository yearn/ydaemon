package controllers

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/majorfi/ydaemon/internal/utils"
	"github.com/sirupsen/logrus"
)

type controller struct{}

func valueWithFallback(s string, defaultValue string) string {
	if s == "" || s == "\"\"" {
		return defaultValue
	}
	return s
}

func bValueWithFallbackUint64(s *big.Int, defaultValue uint64) uint64 {
	if s == nil {
		return defaultValue
	}
	return s.Uint64()
}
func bValueWithFallbackString(s *big.Int, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return s.String()
}

func strToUint(s string, defaultValue uint64) uint64 {
	if s == "" {
		return defaultValue
	}
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return value
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
