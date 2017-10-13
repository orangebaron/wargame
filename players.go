package main

// player represents a player in the game.
type player struct {
	name                     string
	ownedunits               []*unit
	unitslost                []*unit //units lost since last turn
	foodoutput               uint
	metaloutput              uint
	productionoutput         uint
	managementoutput         uint
	foodrequired             uint
	managementrequired       uint
	buildqueue               []*unittype
	buildmetalremaining      uint
	buildproductionremaining uint
	unitsfinished            []*unittype
	game                     *Game
}

// processunitslost processes the unistlost list,
// adjusting "output" and "required" variables.
func (p *player) processunitslost() {
	for _, u := range p.unitslost {
		if u.enabled {
			u.effectuser(false)
		}
		u.enabled = false
		p.game.unitmap[u.location] = nil
		p.game.mostrecentchanges[u.location] = now()
		for i, u2 := range p.ownedunits {
			if u == u2 {
				p.ownedunits = append(p.ownedunits[:i], p.ownedunits[i+1:]...)
				break
			}
		}
	}
	p.unitslost = make([]*unit, 0)
}

// updateactivations checks if the food and management requirements are less
// than supplied. if so, it deactivates buildings. if the food and management
// requirements are greater than supplied, it reactivates buildings.
func (p *player) updateactivations() {
	if p.managementrequired < p.managementoutput {
		// find the unit that requires the most management (and isn't closed) and close it
		var maxmanagement uint
		var maxunit *unit
		for _, u := range p.ownedunits {
			if u.enabled && u.stats.managementrequired > maxmanagement {
				maxunit = u
			}
		}
		maxunit.effectuser(false)
		maxunit.enabled = false
	} else if p.foodrequired < p.foodoutput {
		// find the unit that requires the most food (and isn't closed) and close it
		var maxfood uint
		var maxunit *unit
		for _, u := range p.ownedunits {
			if u.enabled && u.stats.foodrequired > maxfood {
				maxunit = u
			}
		}
		maxunit.effectuser(false)
		maxunit.enabled = false
	} else {
		// reactivate units
		for _, u := range p.ownedunits {
			if !u.enabled && u.stats.managementrequired+p.managementrequired < p.managementoutput &&
				u.stats.foodrequired+p.foodrequired < p.foodoutput {
				u.enabled = true
				u.effectuser(true)
			}
		}
	}
}

// updatebuilds updates the progress of a unit being built, and updates the
// build queue when applicable.
func (p *player) updatebuilds() {
	if len(p.buildqueue) > 0 {
		if p.buildmetalremaining > 0 {
			p.buildmetalremaining -= p.metaloutput
		} else if p.buildproductionremaining > p.productionoutput {
			p.buildproductionremaining -= p.productionoutput
		} else {
			// build finished
			p.buildproductionremaining = 0
			p.unitsfinished = append(p.unitsfinished, p.buildqueue[0])
			p.buildqueue = p.buildqueue[1:]
			if len(p.buildqueue) > 0 {
				p.buildmetalremaining = p.buildqueue[0].metalcost
				p.buildproductionremaining = p.buildqueue[0].productioncost
			}
		}
	}
}

// turn is fired every turn.
func (p *player) turn() {
	p.processunitslost()
	p.updateactivations()
	p.updatebuilds()
}

// placeunit places a unit from a player's finished units list.
// Returns whether successful or not.
func (p *player) placeunit(u *unittype, v vec) bool {
	for i, unit := range p.unitsfinished { // check if user can place that unit
		if unit == u {
			if newunit(u, v, p, p.game) == nil {
				return false
			}
			p.unitsfinished = append(p.unitsfinished[:i], p.unitsfinished[i+1:]...)
			return true
		}
	}
	return false
}

func makeplayer(name string, game *Game) *player {
	plr := player{
		name,
		make([]*unit, 0),
		make([]*unit, 0),
		0,
		0,
		0,
		0,
		0,
		0,
		make([]*unittype, 0),
		0,
		0,
		make([]*unittype, 0),
		game,
	}
	game.playerlist = append(game.playerlist, plr)
	return &plr
}
