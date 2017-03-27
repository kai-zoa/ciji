package battery

func New() Battery {
	return &nilBattery{}
}

type nilBattery struct{}

func (b *nilBattery) RemainingCapacities() int {
	return 0
}
