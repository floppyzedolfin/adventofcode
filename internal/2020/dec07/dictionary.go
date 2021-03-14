package dec07

import (
	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

// bagDictionary lets you find a bag based on its name.
type bagDictionary map[bagColour]*bag

func countUniqueAncestorColours(d bagDictionary, colour string) int {
	ancestors := unique(d.getAllAncestors(bagColour(colour)))
	return len(ancestors)
}

func (d bagDictionary) getAllAncestors(colour bagColour) []bagColour {
	ancestors := make([]bagColour, 0)
	for c := range d[colour].parents {
		ancestors = append(ancestors, c)
		// this is where we hope there is no loop
		ancestors = append(ancestors, d.getAllAncestors(c)...)
	}
	return ancestors
}

// registerBag adds a bag to the dictionary
// loops are currently fully allowed
func (d *bagDictionary) registerBag(b bag) {
	if _, ok := (*d)[b.colour]; !ok {
		// first case scenario - the bag is unknown
		(*d)[b.colour] = &b
	} else {
		// second option - we already know of the bag, and need to update it
		(*d)[b.colour].includeChildren(b)
	}
	// update children and give them a parent
	for childName, count := range b.children {
		if _, ok := (*d)[childName]; !ok {
			// first time we face this child
			(*d)[childName] = &bag{colour:  childName, children: make(map[bagColour]int), parents: make(map[bagColour]int)}
		}
		// add a parent to the child, with count equal to the bag's definition. Since we're registering this parent, we can assume it wasn't there earlier.
		(*d)[childName].parents[b.colour] = count
	}
}

func countContainedBags(d bagDictionary, colour string) int {
	// need to remove 1, because the final bag isn't included in the total count of subbags.
	return d.countBagChildren(bagColour(colour))-1
}

func (d *bagDictionary) countBagChildren(colour bagColour) int {
	if (*d)[colour].childrenCount != nil {
		return *(*d)[colour].childrenCount
	}
	// this bag counts, let's not forget it.
	total := 1
	for child, count := range (*d)[colour].children {
		total += count *  d.countBagChildren(child)
	}
	(*d)[colour].childrenCount = ptr.Int(total)
	return total
}
