package moon

import (
	"testing"
	"time"
)

func TestMoon_Age(t *testing.T) {
	var tm time.Time
	m := &Moon{}
	tm, _ = time.Parse("20060102", "20170327")
	if e, a := 29, m.Age(tm); e != a {
		t.Errorf("{\n - %v\n + %v\n}", e, a)
	}
	tm, _ = time.Parse("20060102", "20170328")
	if e, a := 0, m.Age(tm); e != a {
		t.Errorf("{\n - %v\n + %v\n}", e, a)
	}
}
