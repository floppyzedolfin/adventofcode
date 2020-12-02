package dec02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// password holds the information about a password: the value, and its policy
type password struct {
	value  string
	policy policy
}

// policy holds the security details that must be enforced
type policy struct {
	min    int
	max    int
	letter string
}

// isValid1 returns whether a password's value matches its policy for Part 1
func (p password) isValid1() bool {
	count := strings.Count(p.value, p.policy.letter)
	return p.policy.min <= count && count <= p.policy.max
}

// isValid2 returns whether a password's value matches its policy for Part 2
func (p password) isValid2() bool {
	if p.policy.min > len(p.value) || p.policy.max > len(p.value) {
		// let's hope this doesn't happen, but also make sure we wouldn't break something ere Christmas.
		return false
	}
	// those -1 are here because someone decided to start counting at 1
	firstPos := string(p.value[p.policy.min-1]) == p.policy.letter
	secondPos := string(p.value[p.policy.max-1]) == p.policy.letter
	// fancy XOR operator
	return firstPos != secondPos
}

const (
	validLineRegExp = `(\d+)-(\d+) ([[:alpha:]]): ([[:alnum:]]+)$`
)

// buildPassword builds a password
func buildPassword(line string) (password, error) {
	// parse the line to extract the information we need
	r := regexp.MustCompile(validLineRegExp)
	res := r.FindStringSubmatch(line)
	if len(res) != 5 {
		return password{}, fmt.Errorf("unable to extract password from line %s", line)
	}

	// convert the values to integers for later use
	min, err := strconv.Atoi(res[1])
	if err != nil {
		return password{}, fmt.Errorf("unable to build password, min %s is not an integer", res[1])
	}
	max, err := strconv.Atoi(res[2])
	if err != nil {
		return password{}, fmt.Errorf("unable to build password, max %s is not an integer", res[2])
	}

	// finally build the password
	return password{
		value:  res[4],
		policy: policy{
			min:    min,
			max:    max,
			letter: res[3],
		},
	}, nil
}
