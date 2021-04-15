package position

// Position holds the coordinates of a point on a map
type Position struct {
	X int
	Y int
}

// L1 returns the L1-distance (sum of norms of coordinates)
func (p *Position) L1() int {
	return abs(p.X) + abs(p.Y)
}

// Move translates the position by the given vector
func (p *Position) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

// Rotate rotates the point around the origin, by a trigonometrical (anti-clockwise) angle.
func (p *Position) Rotate(angle int) {
	switch angle {
	case 90:
		p.X, p.Y = -p.Y, p.X
	case 180:
		p.X, p.Y = -p.X, -p.Y
	case 270:
		p.X, p.Y = p.Y, -p.X
	default:
		// ignore the command
	}
}

// abs returns the absolute value of an int
func abs(i int) int {
	if i <= 0 {
		return -i
	}
	return i
}
