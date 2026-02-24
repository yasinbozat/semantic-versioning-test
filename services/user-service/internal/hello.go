package internal

func Add(a, b int) int {
	return a + b
}

func Discount(price float64, percent float64) float64 {
	return price - (price * percent / 100)
}
