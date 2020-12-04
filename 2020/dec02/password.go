package dec02

import (
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

func countValidPasswordsPrima(passwords []password) int {
	count := 0
	for _, p := range passwords {
		if p.isValid1() {
			count++
		}
	}
	return count
}

func countValidPasswordsSecunda(passwords []password) int {
	count := 0
	for _, p := range passwords {
		if p.isValid2() {
			count++
		}
	}
	return count
}
