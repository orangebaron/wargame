package main

// player represents a player in the game.
type player struct {
	ownedunits  []*unit
	unitslost   []*unit //units lost since last turn
	foodO       uint
	metalO      uint
	productionO uint
	managementO uint
	people      uint
	peopleR     uint
	managementR uint
}

// turn is fired every turn.
func (p *player) turn() {
	for _, u := range p.unitslost {
		p.foodO -= u.stats.foodO
		p.metalO -= u.stats.metalO
		p.productionO -= u.stats.productionO
		p.managementO -= u.stats.managementO
		p.peopleR -= u.stats.peopleR
		p.managementR -= u.stats.managementR
	}
}
