package tokens

import (
	"github.com/yearn/ydaemon/internal/meta"
)

type Controller struct{}

func selectLocalizationFromString(s string, loc meta.TLocalization) meta.TLocalizationDetails {
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
