package dec07

import (
	"fmt"
	"regexp"
	"strconv"
)

// bagColour holds the colour of a bag, and is also an identifier for a bag.
type bagColour string

// bag contains the information of what can be found inside a bag.
type bag struct {
	children map[bagColour]int
	parents  map[bagColour]int
	colour   bagColour
}

// bagDictionary lets you find a bag based on its name.
type bagDictionary map[bagColour]*bag

func countUniqueAncestorColours(d bagDictionary, colour string) int {
	ancestors := unique(d.getAllAncestors(bagColour(colour)))
	return len(ancestors)
}

// unique keeps unique elements from a slice
func unique(colours []bagColour) []bagColour {
	keys := make(map[bagColour]bool)
	uColours := make([]bagColour, 0)
	for _, colour := range colours {
		if _, v := keys[colour]; !v {
			keys[colour] = true
			uColours = append(uColours, colour)
		}
	}
	return uColours
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

func buildBag(colour string, contents []string) (bag, error) {
	b := bag{colour: bagColour(colour), children: make(map[bagColour]int), parents: make(map[bagColour]int)}
	// check if this is a leaf
	const leafRegExp = `no other bags`
	reLeaf := regexp.MustCompile(leafRegExp)
	if len(contents) == 1 && reLeaf.MatchString(contents[0]) {
		return b, nil
	}

	// parse the contents to determine how exactly we need to fill the bag
	const childInfoRegExp = `\s*([0-9]+) ([a-z ]+) bags?\s*`
	re := regexp.MustCompile(childInfoRegExp)
	for _, c := range contents {
		res := re.FindStringSubmatch(c)
		if len(res) != 3 {
			// missing something here
			return bag{}, fmt.Errorf("invalid bag contents: '%s'", c)
		}
		childColour := bagColour(res[2])
		childCount, err := strconv.Atoi(res[1])
		if err != nil {
			return bag{}, fmt.Errorf("invalid field child count: '%s'", c)
		}
		b.children[childColour] = childCount
	}
	return b, nil
}

func (b *bag) includeChildren(other bag) {
	for colour, count := range other.children {
		if c, ok := b.children[colour]; !ok {
			b.children[colour] = count
		} else {
			b.children[colour] = c + count
		}
	}
}
