package core

// Game represents a game and includes its data.
type Game struct {
	// UnitMap is a map of a location on the board to the unit in that location.
	UnitMap map[Vec]*Unit
	// MostRecentChanges list the most recent time a spot has been changed.
	MostRecentChanges map[Vec]uint
	// Playerlist is a list of players in the game.
	PlayerList     []Player
	CurrentTurnNum uint
}

// MakeGame makes a new Game with default parameters.
func MakeGame() *Game {
	return &Game{
		make(map[Vec]*Unit),
		make(map[Vec]uint),
		make([]Player, 0),
		0,
	}
}
