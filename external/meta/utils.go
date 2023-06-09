package meta

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yearn/ydaemon/internal/models"
)

type Controller struct{}

func selectLocalizationFromString(s string, loc models.TLocalization) models.TLocalizationDetails {
	switch s {
	case `en`:
		return loc.En
	case `fr`:
		return loc.Fr
	case `es`:
		return loc.Es
	case `de`:
		return loc.De
	case `pt`:
		return loc.Pt
	case `el`:
		return loc.El
	case `tr`:
		return loc.Tr
	case `vi`:
		return loc.Vi
	case `zh`:
		return loc.Zh
	case `hi`:
		return loc.Hi
	case `ja`:
		return loc.Ja
	case `id`:
		return loc.Id
	case `ru`:
		return loc.Ru
	}
	return loc.En
}

func getQuery(c *gin.Context, targetKey string) string {
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if strings.EqualFold(targetKey, key) {
			return strings.Join(values, ",")
		}
	}
	return ""
}
