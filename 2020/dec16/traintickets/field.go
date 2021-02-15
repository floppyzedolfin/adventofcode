package traintickets

type field struct {
	name string
	ranges []minMax

	// possibleEntries holds the list of entries for which this field could be valid
	possibleEntries map[int]struct{}
}

type minMax struct {
	min int
	max int
}

// sieve scans the ticket and removes fields from the possible entries if they're invalid
func (f *field) sieve(t ticket) {
	for i, v := range t.values {
		removeEntry := true
		for _, r := range f.ranges {
			if r.min <=  v && v <= r.max {
				removeEntry = false
			}
		}
		if removeEntry {
			delete(f.possibleEntries, i)
		}
	}
}
