package commodity

type Product struct {
	Name string
	Type int
	Cost float32
}

func (p *Product) CalculatePrice() float32 {
	return p.Cost
}
