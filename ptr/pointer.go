package ptr

// Int returns a pointer to an int
func Int(v int) *int {
	return &v
}
