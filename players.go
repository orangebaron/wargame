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
		if u.enabled {
			u.effectuser(false)
		}
		u.enabled = false
		unitmap[u.location] = nil
		mostrecentchanges[u.location] = now()
		for i, u2 := range p.ownedunits {
			if u == u2 {
				p.ownedunits = append(p.ownedunits[:i], p.ownedunits[i+1:]...)
				break
			}
		}
	}
	p.unitslost = make([]*unit, 0)

	// change population
	const growthrate = 10
	if p.foodO > p.people {
		p.people += 1 + ((p.foodO - p.people) / 10)
	} else {
		p.people -= 1 + ((p.people - p.foodO) / 10)
	}

	// check if people and management requirements are met or if any closed buildings can be opened
	if p.managementR < p.managementO {
		// find the unit that requires the most management (and isn't closed) and close it
		var maxmanagement uint
		var maxunit *unit
		for _, u := range p.ownedunits {
			if u.enabled && u.stats.managementR > maxmanagement {
				maxunit = u
			}
		}
		maxunit.effectuser(false)
		maxunit.enabled = false
	} else if p.peopleR < p.people {
		// find the unit that requires the most people (and isn't closed) and close it
		var maxpeople uint
		var maxunit *unit
		for _, u := range p.ownedunits {
			if u.enabled && u.stats.peopleR > maxpeople {
				maxunit = u
			}
		}
		maxunit.effectuser(false)
		maxunit.enabled = false
	} else {
		// reactivate units
		for _, u := range p.ownedunits {
			if !u.enabled && u.stats.managementR+p.managementR < p.managementO && u.stats.peopleR+p.peopleR < p.people {
				u.enabled = true
				u.effectuser(true)
			}
		}
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
