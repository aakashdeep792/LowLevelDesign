package commodity

type flatDiscountCoupon struct {
	item     Commodity
	discount float32
	minPrice float32
	pType    int
}

func NewFlatDiscountCoupon(item Commodity, pType int, discount, minPrice float32) Commodity {
	return &flatDiscountCoupon{
		item:     item,
		discount: discount,
		minPrice: minPrice,
		pType:    pType,
	}
}

func (p *flatDiscountCoupon) isValidTypes() bool {
	eligibleItem := []int{ELECTRONIC, FURNITURE}

	for _, v := range eligibleItem {
		if v == p.pType {
			return true
		}
	}

	return false
}

func (p *flatDiscountCoupon) CalculatePrice() float32 {
	// it will be applied to all the product
	price := p.item.CalculatePrice()

	if price < p.minPrice || !p.isValidTypes() {
		return price
	}

	newPrice := price - p.discount

	return newPrice
}
