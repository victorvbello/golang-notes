package shakes

type ChocolateTopping struct {
	shake Shake
}

func NewChocolateTopping(shake Shake) *ChocolateTopping {
	return &ChocolateTopping{shake: shake}
}

func (c *ChocolateTopping) GetDescription() string {
	return c.shake.GetDescription() + ", Chocolate Topping"
}

func (c *ChocolateTopping) Cost() float64 {
	return c.shake.Cost() + 1.0
}

type StrawberryTopping struct {
	shake Shake
}

func NewStrawberryTopping(shake Shake) *StrawberryTopping {
	return &StrawberryTopping{shake: shake}
}

func (s *StrawberryTopping) GetDescription() string {
	return s.shake.GetDescription() + ", Strawberry Topping"
}

func (s *StrawberryTopping) Cost() float64 {
	return s.shake.Cost() + 1.5
}

type DonutTopping struct {
	shake Shake
}

func NewDonutTopping(shake Shake) *DonutTopping {
	return &DonutTopping{shake: shake}
}

func (d *DonutTopping) GetDescription() string {
	return d.shake.GetDescription() + ", Donut Topping"
}

func (d *DonutTopping) Cost() float64 {
	return d.shake.Cost() + 2.0
}
