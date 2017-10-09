package main

// player represents a player in the game.
type player struct {
	ownedunits               []*unit
	unitslost                []*unit //units lost since last turn
	foodO                    uint
	metalO                   uint
	productionO              uint
	managementO              uint
	people                   uint
	peopleR                  uint
	managementR              uint
	buildqueue               []*unittype
	buildmetalremaining      uint
	buildproductionremaining uint
	unitsfinished            []*unittype
}

// turn is fired every turn.
func (p *player) turn() {
	// change resources outputs and required resources
	for _, u := range p.unitslost {
		p.foodO -= u.stats.foodO
		p.metalO -= u.stats.metalO
		p.productionO -= u.stats.productionO
		p.managementO -= u.stats.managementO
		p.peopleR -= u.stats.peopleR
		p.managementR -= u.stats.managementR
	}
	// change what's being built
	if len(p.buildqueue) > 0 {
		if p.buildmetalremaining > 0 {
			p.buildmetalremaining -= p.metalO
		} else if p.buildproductionremaining > p.productionO {
			p.buildproductionremaining -= p.productionO
		} else {
			// build finished
			p.buildproductionremaining = 0
			p.unitsfinished = append(p.unitsfinished, p.buildqueue[0])
			p.buildqueue = p.buildqueue[1:]
			if len(p.buildqueue) > 0 {
				p.buildmetalremaining = p.buildqueue[0].metalC
				p.buildproductionremaining = p.buildqueue[0].productionC
			}
		}
	}
}

// placeunit places a unit from a player's finished units list.
// Returns whether successful or not.
func (p *player) placeunit(u *unittype, v vec) bool {
	for i, unit := range p.unitsfinished { // check if user can place that unit
		if unit == u {
			if newunit(u, v, p) == nil {
				return false
			}
			p.unitsfinished = append(p.unitsfinished[:i], p.unitsfinished[i+1:]...)
			return true
		}
	}
	return false
}
