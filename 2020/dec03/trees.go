package dec03

const tree = "#"

func countTreesPart1(slope []string) int {
	return countTrees(slope, 3,1 )
}

func countTrees(slope []string, sleighX, sleighY int ) int {
	xPosition := 0
	treesHit := 0
	if len(slope) == 0 {
		// no slope, no trees, no hit.
		return 0
	}
	// we can assume, from the exercise, that each line has the same width.
	trackWidth := len(slope[0])
	if trackWidth == 0 {
		// we're not jumping off the cliff to get to the bottom of the mountain.
		return 0
	}

	for travelled := 0; travelled < len(slope); travelled += sleighY {
		// did we hit a tree ?
		if string(slope[travelled][xPosition]) == tree {
			treesHit++
		}
		// advance horizontally
		xPosition = (xPosition + sleighX) % trackWidth
	}
	return treesHit
}

func countTreesPart2(slope []string) int {
	check1 := countTrees(slope, 1, 1)
	check2 := countTrees(slope, 3, 1)
	check3 := countTrees(slope, 5, 1)
	check4 := countTrees(slope, 7, 1)
	check5 := countTrees(slope, 1, 2)
	return check1*check2*check3*check4*check5
}

