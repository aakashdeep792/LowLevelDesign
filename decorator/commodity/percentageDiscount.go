package commodity

type percentageDiscountCoupon struct {
	item     Commodity
	discount float32
}

func NewPercentageDiscountCoupon(item Commodity, discount float32) Commodity {
	return &percentageDiscountCoupon{
		item:     item,
		discount: discount,
	}
}

func (p *percentageDiscountCoupon) CalculatePrice() float32 {
	// it will be applied to all the product
	price := p.item.CalculatePrice()
	newPrice := (100 - p.discount) * price / 100

	return newPrice
}
