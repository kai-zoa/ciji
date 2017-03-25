package moon

import (
	"time"
)

const (
	MaxAge = 29
)

type Moon struct{}

func (*Moon) Age(now time.Time) int {
	y, m, d := now.Date()
	g := ((((y-2009)%19)*11 + int(m) + d) + 1)
	if int(m) < 3 {
		g = g + 2
	}
	g = g % 30
	return g
}
