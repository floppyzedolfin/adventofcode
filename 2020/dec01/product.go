package dec01

// product performs the product of ints in a slice.
func product(data []int) int {
	p := 1
	for _, d := range data {
		p *= d
	}
	return p
}

