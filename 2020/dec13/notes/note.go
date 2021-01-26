package notes

// Notes is used to parse the input file. It maps its contents to a usable structure
type Notes struct {
	initialised bool // we've read the first line
	currentTime int
	buses       []bus
}
