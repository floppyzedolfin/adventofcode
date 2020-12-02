package door

// Solver defines what you can expect from an exercise
type Solver interface {
	Solve() (Result, error)
}
