package instructions

// Action describes the move the ship will operate
type Action string

const (
	MoveNorth     Action = "N"
	MoveEast      Action = "E"
	MoveSouth     Action = "S"
	MoveWest      Action = "W"
	TurnPort      Action = "L"
	TurnStarboard Action = "R"
	Forward       Action = "F"
)
