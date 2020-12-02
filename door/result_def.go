package door

// Result is what you could expect from a successful run
type Result interface {
	String() string
}