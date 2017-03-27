package wifi

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

const (
	cmdAirport = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
)

func New() WIFI {
	return &macOSAirport{}
}

type macOSAirport struct {
	data map[string]string
}

func (w *macOSAirport) SSID() string {
	w.update()
	return w.data["SSID"]
}

func (w *macOSAirport) Intensity() float64 {
	w.update()
	rssi := w.data["agrCtlRSSI"]
	n, err := strconv.Atoi(rssi)
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
	// RSSI
	// -20  Excellent
	// -30  Excellent
	// -40  Excellent
	// -50  Excellent
	// -60  better
	// -70  good
	// -80  not good
	// -90  bad
	// -100 bad
	if n < -80 {
		return 0.2
	}
	if n < -60 {
		return 0.6
	}
	return 1.0
}

func (w *macOSAirport) update() {
	if len(w.data) > 0 {
		return
	}
	var err error
	out, err := exec.Command(cmdAirport, "-I").Output()
	if err != nil {
		panic(errors.Wrap(err, ""))
	}

	w.data = make(map[string]string)

	r := bufio.NewReader(bytes.NewReader(out))
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(errors.Wrap(err, ""))
		}
		entry := strings.Split(string(line), ":")
		if len(entry) > 1 {
			k, v := strings.Trim(entry[0], " "), strings.Trim(entry[1], " ")
			w.data[k] = v
		}
	}
}
