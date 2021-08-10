package encapsulamiento

type auto struct {
	model string
	year  int
	name  string
}

func NewAuto() *auto {
	return &auto{}
}

func (auto *auto) SetName(name string) {
	auto.name = name
}

func (auto *auto) GetName() string {
	return auto.name
}
