package cart

import "LowLevelDesign/decorator/commodity"

type shoppingCart struct {
	items []commodity.Commodity
}

func NewCart() *shoppingCart {

	return &shoppingCart{
		items: make([]commodity.Commodity, 0),
	}

}
func (c *shoppingCart) AddToCart(item *commodity.Product) {
	tmp := commodity.NewPercentageDiscountCoupon(commodity.NewFlatDiscountCoupon(item, item.Type, 300, 3000), 10)

	c.items = append(c.items, tmp)
}

func (c *shoppingCart) CalculateTotalPrice() float32 {
	var total float32
	for _, v := range c.items {
		total += v.CalculatePrice()
	}

	return total
}
