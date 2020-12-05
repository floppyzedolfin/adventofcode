package dec03

import (
	"fmt"

	"github.com/floppyzedolfin/adventofcode/fileparser"
)

// readForest reads a whole file into memory
// and returns a forest - its lines.
func readForest(path string) (forest, error) {
	var f forest
	err := fileparser.ParseFile(path, &f)
	if err != nil {
		return forest{}, fmt.Errorf("unable to parse file '%s': %s", path, err.Error())
	}
	return f, nil
}

func (f *forest) ParseLine(line string) error {
	f.environment = append(f.environment, []byte(line))
	return nil
}
