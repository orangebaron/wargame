package main

// Game represents a game and includes its data.
type Game struct {
	// unitmap is a map of a location on the board to the unit in that location.
	unitmap map[vec]*unit
	// mostrecentchanges list the most recent time a spot has been changed.
	mostrecentchanges map[vec]time
	// playerlist is a list of players in the game.
	playerlist []player
}

// MakeGame makes a new Game with default parameters.
func MakeGame() *Game {
	return &Game{
		make(map[vec]*unit),
		make(map[vec]time),
		make([]player, 0),
	}
}
