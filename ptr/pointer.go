package ptr

// Int returns a pointer to an int
func Int(v int) *int {
	return &v
}

// Int64 returns a pointer to an int64
func Int64(v int64) *int64 {
	return &v
}
