package indicator

type Source interface {
	Progress() float64
}

type Indicator interface {
	String() string
}
