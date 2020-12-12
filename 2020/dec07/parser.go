package dec07

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/floppyzedolfin/adventofcode/fileparser"
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
