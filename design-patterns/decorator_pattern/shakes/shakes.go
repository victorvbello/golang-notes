package shakes

type Shake interface {
	GetDescription() string
	Cost() float64
}

type MilkShake struct{}

func (m *MilkShake) GetDescription() string {
	return "Milk Shake"
}

func (m *MilkShake) Cost() float64 {
	return 5.0
}

type ChocolateShake struct{}

func (c *ChocolateShake) GetDescription() string {
	return "Chocolate Shake"
}

func (c *ChocolateShake) Cost() float64 {
	return 6.0
}
