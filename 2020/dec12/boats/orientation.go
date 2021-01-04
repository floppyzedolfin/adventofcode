package boats

// cleanOrientation guarantees orientation is between 0 and 359
func cleanOrientation(orientation int) int {

	// how often do you get to use this fancy operator?
	orientation %= 360

	if orientation < 0 {
		orientation += 360
	}

	return orientation
}
