package indicator

import (
	"fmt"
)

type percentageIndicator struct {
	src Source
}

func (n *percentageIndicator) String() string {
	pct := int(n.src.Progress() * 100)
	if pct > 100 {
		pct = 100
	}
	return fmt.Sprint(pct)
}

func NewPercentage(s Source) Indicator {
	return &percentageIndicator{
		src: s,
	}
}
