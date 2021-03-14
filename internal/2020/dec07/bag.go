package dec07

// bagColour holds the colour of a bag, and is also an identifier for a bag.
type bagColour string

// bag contains the information of what can be found inside a bag, and of where we can find bags.
type bag struct {
	children map[bagColour]int
	parents  map[bagColour]int
	colour   bagColour
	childrenCount *int
}

// unique keeps unique elements from a slice
func unique(colours []bagColour) []bagColour {
	keys := make(map[bagColour]bool)
	uColours := make([]bagColour, 0)
	for _, colour := range colours {
		if v := keys[colour]; !v {
			keys[colour] = true
			uColours = append(uColours, colour)
		}
	}
	return uColours
}


// includeChildren includes the children of a second bag to the current bag. Sort of like "adopting" a bag's kids.
func (b *bag) includeChildren(other bag) {
	for colour, count := range other.children {
		if c, ok := b.children[colour]; !ok {
			b.children[colour] = count
		} else {
			b.children[colour] = c + count
		}
	}
}
