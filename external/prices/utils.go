package prices

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func getQuery(c *gin.Context, targetKey string) string {
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if strings.EqualFold(targetKey, key) {
			return strings.Join(values, ",")
		}
	}
	return ""
}
