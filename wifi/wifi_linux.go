package wifi

func New() WIFI {
	return &nilWIFI{}
}

type nilWIFI struct{}

func (w *nilWIFI) SSID() string {
	return ""
}

func (w *nilWIFI) Intensity() float64 {
	return 0.0
}
