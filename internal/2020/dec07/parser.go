package dec07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/fileparser"
)

func readBags(inputPath string) (bagDictionary, error) {
	dic := make(bagDictionary)
	err := fileparser.ParseFile(inputPath, &dic)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file '%s': %s", inputPath, err.Error())
	}
	return dic, nil
}

// ParseLine converts a string to something we'll add in our dictionary
func (d *bagDictionary) ParseLine(line string) error {
	// a line consists of
	// - the container: 'shiny gold bags'
	// - the word 'contain'
	// - a repeated section that contains
	//   - a value: '4'
	//   - a containee: 'faded blue bags' (which might or mightn't end with an s, according to grammar laws)
	//   - a comma or a full stop
	// for clarity, I'll parse the contained things in a second part
	const validLineRegExp = `([a-z ]+) bags contain ([a-z0-9, ]+).$`

	// parse the line to extract the information we need
	r := regexp.MustCompile(validLineRegExp)
	res := r.FindStringSubmatch(line)
	if len(res) <= 2 {
		// something is wrong here, the line is prabably not valid.
		return fmt.Errorf("unable to parse line '%s'", line)
	}

	// build this bag
	b, err := buildBag(res[1], strings.Split(res[2], ","))
	if err != nil {
		return fmt.Errorf("unable to build bag from line '%s': %s", line, err.Error())
	}
	d.registerBag(b)
	return nil
}

// buildBag creates a bag from the contents we read from a line
func buildBag(colour string, contents []string) (bag, error) {
	b := bag{colour: bagColour(colour), children: make(map[bagColour]int), parents: make(map[bagColour]int)}
	// check if this is a leaf
	const leaf = "no other bags"
	if len(contents) == 1 && contents[0] == leaf {
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
