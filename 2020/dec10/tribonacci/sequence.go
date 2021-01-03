package tribonacci

// Sequence is an object that holds data of a Tribonacci sequence (s[n+1] = s[n] + s[n-1] + s[n-2])
type Sequence struct {
	values []int64
}


// New returns a Tribonacci sequence initialized with the input values
func New(n0, n1, n2 int64) *Sequence {
	return &Sequence{
		// fill with non-0 data
		values: []int64{n0, n1, n2},
	}
}

// At returns the "Tribonacci" value at index n
func (s *Sequence) At(n int) int64 {
	if n >= len(s.values) {
		// memoisation
		s.buildTillIndex(n)
	}
	return s.values[n]
}

// buildTillIndex completes the Tribonacci sequence until it has the requested index
func (s *Sequence) buildTillIndex(n int) {
	tribLength := len(s.values)
	for t := tribLength; t <= n; t++ {
		s.values = append(s.values, s.values[t-1]+s.values[t-2]+s.values[t-3])
	}
}
