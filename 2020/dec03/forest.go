package dec03

const tree = '#'

func (f forest) countTreesPrima() int {
	return f.countTrees(3,1)
}

func (f forest) countTrees(sleighX, sleighY int) int {
	xPosition := 0
	treesHit := 0
	if len(f.environment) == 0 {
		// no slope, no trees, no hit.
		return 0
	}
	// we can assume, from the exercise, that each line has the same width.
	trackWidth := len(f.environment[0])
	if trackWidth == 0 {
		// we're not jumping off the cliff to get to the bottom of the mountain.
		return 0
	}

	for travelled := 0; travelled < len(f.environment); travelled += sleighY {
		// did we hit a tree ?
		if f.environment[travelled][xPosition] == tree {
			treesHit++
		}
		// advance horizontally
		xPosition = (xPosition + sleighX) % trackWidth
	}
	return treesHit
}

func (f forest) countTreesSecunda() int {
	check1 := f.countTrees(1, 1)
	check2 := f.countTrees(3, 1)
	check3 := f.countTrees(5, 1)
	check4 := f.countTrees(7, 1)
	check5 := f.countTrees(1, 2)
	return check1*check2*check3*check4*check5
}

// forest holds the list of stuff we can find on the mountain
type forest struct {
	environment [][]byte
}
