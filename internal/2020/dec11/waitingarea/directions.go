package waitingarea

var (
	// directions holds the list of vectors whereupon the passengers will lay their eyes
	directions = []vector{
		{-1, -1}, // In the world of advertising, there's no such thing as a lie. There's only expedient exaggeration.
		{0, -1},  // due north
		{1, -1},  // north east
		{-1, 0},  // life is peaceful there // in the open air
		{1, 0},   // due east
		{-1, 1},  // south west
		{0, 1},   // south
		{1, 1},   // south east
	}
)
