package commodity

const (
	ELECTRONIC int = iota
	CLOTHING
	FOOTWARE
	FURNITURE
)

type Commodity interface {
	CalculatePrice() float32
}
