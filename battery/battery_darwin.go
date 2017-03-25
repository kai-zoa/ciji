package battery

import (
	"github.com/pkg/errors"
	"os/exec"
	"regexp"
	"strconv"
)

func New() Battery {
	return &darwinBattery{}
}

type darwinBattery struct{}

func (b *darwinBattery) RemainingCapacities() int {
	var err error
	out, err := exec.Command("pmset", "-g", "batt").Output()
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
	r := regexp.MustCompile(`([\d]+)%`)
	matches := r.FindAllStringSubmatch(string(out), -1)
	if len(matches) > 0 {
		g := matches[0]
		if len(g) > 1 {
			n, err := strconv.Atoi(g[1])
			if err != nil {
				panic(errors.Wrap(err, ""))
			}
			return n
		}
	}
	return 0
}
