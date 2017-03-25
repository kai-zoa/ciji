package indicator

import (
	"math"
)

type themeIndicator struct {
	src   Source
	theme []string
}

func (n *themeIndicator) String() string {
	progress := n.src.Progress()
	if progress >= 1 {
		progress = 1
	}
	i := int(math.Floor(float64(len(n.theme)) * progress))
	if i < len(n.theme) {
		return n.theme[i]
	}
	return n.theme[len(n.theme)-1]
}

func NewTheme(s Source, theme []string) Indicator {
	return &themeIndicator{
		src:   s,
		theme: theme,
	}
}
