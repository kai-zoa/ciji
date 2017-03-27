package main

import (
	"fmt"
	"github.com/kai-zoa/ciji/battery"
	"github.com/kai-zoa/ciji/indicator"
	"github.com/kai-zoa/ciji/moon"
	"github.com/kai-zoa/ciji/wifi"
	"os"
	"text/template"
	"time"
)

func usage() {
	fmt.Println(`Usage: ciji "[text tempate]"`)
}

type NodeAttr interface {
	String() string
}

type Node map[string]NodeAttr
type NodeList map[string]Node

type batterySource struct {
	batt battery.Battery
}

func (s *batterySource) Progress() float64 {
	cap := s.batt.RemainingCapacities()
	return float64(cap) / 100.0
}

type wifiIntensitySource struct {
	w wifi.WIFI
}

func (s *wifiIntensitySource) Progress() float64 {
	return s.w.Intensity()
}

type wifiSSID struct {
	w wifi.WIFI
}

func (s wifiSSID) String() string {
	return s.w.SSID()
}

type moonSource struct {
	now time.Time
	m   moon.Moon
}

func (s *moonSource) Progress() float64 {
	age := s.m.Age(s.now)
	return float64(age) / float64(moon.MaxAge)
}

func main() {

	args := os.Args
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	srcBatt := &batterySource{
		batt: battery.New(),
	}

	srcMoon := &moonSource{
		now: time.Now(),
	}

	wifiObj := wifi.New()

	srcWIFI := &wifiIntensitySource{
		w: wifiObj,
	}

	tmpl := args[1]
	t := template.New("ciji")
	_, err := t.Parse(tmpl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "template format error: \"%s\"\n%+v", tmpl, err)
		os.Exit(1)
	}

	t.Execute(os.Stdout, NodeList{
		"Battery": Node{
			"Percentage":   indicator.NewPercentage(srcBatt),
			"VProgressBar": indicator.NewTheme(srcBatt, indicator.VProgressBar),
			"HProgressBar": indicator.NewTheme(srcBatt, indicator.HProgressBar),
			"TMUXColor":    indicator.NewTheme(srcBatt, indicator.TMUXRedToGreen),
		},
		"WIFI": Node{
			"Intensity": indicator.NewTheme(srcWIFI, indicator.Signals),
			"SSID":      &wifiSSID{w: wifiObj},
		},
		"MoonPhase": Node{
			"EMOJI": indicator.NewTheme(srcMoon, indicator.EMOJIMoonPhase),
		},
	})
	os.Stdout.WriteString("\n")
}
