package main

import (
	"LowLevelDesign/decorator/cart"
	"LowLevelDesign/decorator/commodity"
	"fmt"
)

func main() {
	fmt.Println("welcome outstanding programmer aakash")
	c := cart.NewCart()

	table := &commodity.Product{Name: "Nilkamal Table",
		Type: commodity.FURNITURE, Cost: 4000}
	c.AddToCart(table)
	chair := &commodity.Product{Name: "Nilkamal Chair",
		Type: commodity.FURNITURE, Cost: 2000}
	c.AddToCart(chair)
	tv := &commodity.Product{Name: "LG TV",
		Type: commodity.ELECTRONIC, Cost: 40000}
	c.AddToCart(tv)
	shirt := &commodity.Product{Name: "Nilkamal Table",
		Type: commodity.FURNITURE, Cost: 4000}
	c.AddToCart(shirt)

	fmt.Println("total item price ", c.CalculateTotalPrice())
}
