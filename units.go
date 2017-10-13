package main

// unittype contains data for a type of unit.
// a unit can be a boat or a tile.
type unittype struct {
	name               string
	foodoutput         uint
	metaloutput        uint
	productionoutput   uint
	managementoutput   uint
	peoplerequired     uint
	managementrequired uint
	metalcost          uint
	productioncost     uint
	maxhealth          uint
	speed              uint
	hitrange           uint
	damage             uint
}

// types is a list of all possible unit types.
var types = [14]unittype{
	unittype{"Farm", 75, 0, 0, 0, 5, 1, 1, 20, 3, 0, 0, 0},
	unittype{"Urban Farm", 200, 0, 0, 0, 10, 3, 5, 50, 1, 0, 0, 0},
	unittype{"Mineshaft", 0, 1, 0, 0, 30, 7, 3, 15, 3, 0, 0, 0},
	unittype{"Factory", 0, 0, 1, 0, 40, 10, 10, 60, 1, 0, 0, 0},
	unittype{"Robotic Factory", 0, 0, 3, 0, 0, 7, 20, 120, 1, 0, 0, 0},
	unittype{"Small Office", 0, 0, 0, 100, 20, 0, 4, 40, 1, 0, 0, 0},
	unittype{"Office Tower", 0, 0, 0, 300, 50, 0, 10, 100, 1, 0, 0, 0},
	unittype{"Wall", 0, 0, 0, 0, 0, 0, 10, 15, 10, 0, 0, 0},
	unittype{"Turret", 0, 0, 0, 0, 2, 1, 15, 30, 5, 0, 1, 5},
	unittype{"Missile Station", 0, 0, 0, 0, 5, 2, 20, 40, 5, 0, 20, 10},
	unittype{"Cruiser", 0, 0, 0, 0, 10, 4, 40, 90, 5, 5, 5, 7},
	unittype{"Frigate", 0, 0, 0, 0, 7, 3, 20, 50, 4, 7, 2, 4},
	unittype{"Destroyer", 0, 0, 0, 0, 10, 4, 35, 90, 5, 10, 2, 5},
	unittype{"Aircraft Carrier", 0, 0, 0, 0, 50, 6, 50, 120, 15, 5, 10, 10},
}

// vec represents a 2d vector.
type vec struct {
	x int
	y int
}

func (v vec) add(v2 vec) vec {
	return vec{v.x + v2.x, v.y + v2.y}
}
func (v vec) sub(v2 vec) vec {
	return vec{v.x - v2.x, v.y - v2.y}
}
func (v vec) mult(v2 vec) vec {
	return vec{v.x * v2.x, v.y * v2.y}
}
func (v vec) div(v2 vec) vec {
	return vec{v.x / v2.x, v.y / v2.y}
}

// unit represents a unit or tile.
type unit struct {
	stats    *unittype
	health   uint
	location vec
	owner    *player
	enabled  bool
}

// unitmap is a map of a location on the board to the unit in that location.
var unitmap = make(map[vec]*unit)

// mostrecentchanges list the most recent time a spot has been changed.
var mostrecentchanges = make(map[vec]time)

func (u *unit) move(v vec) bool {
	if unitmap[v] == nil && u.enabled {
		u.location = v
		return true
	}
	return false
}

// effectuser applies output and requirement effects to the player's values
func (u *unit) effectuser(positive bool) {
	p := u.owner
	if positive {
		p.foodoutput += u.stats.foodoutput
		p.metaloutput += u.stats.metaloutput
		p.productionoutput += u.stats.productionoutput
		p.managementoutput += u.stats.managementoutput
		p.peoplerequired += u.stats.peoplerequired
		p.managementrequired += u.stats.managementrequired
	} else {
		p.foodoutput -= u.stats.foodoutput
		p.metaloutput -= u.stats.metaloutput
		p.productionoutput -= u.stats.productionoutput
		p.managementoutput -= u.stats.managementoutput
		p.peoplerequired -= u.stats.peoplerequired
		p.managementrequired -= u.stats.managementrequired
	}
}
func (u *unit) damage(d uint) bool { //return value: whether unit was killed
	if u.health <= d {
		u.health = 0
		u.owner.unitslost = append(u.owner.unitslost, u) // register with owner that unit is dead
		return true
	}
	u.health -= d
	return false
}
func (u *unit) attack(target *unit) bool {
	if !u.enabled {
		return false
	}
	return target.damage(u.stats.damage)
}

// newunit creates and initializes a new unit with the given specifications.
func newunit(stats *unittype, location vec, owner *player) *unit {
	if unitmap[location] != nil {
		return nil
	}

	u := unit{stats, stats.maxhealth, location, owner, true}
	unitmap[location] = &u
	owner.ownedunits = append(owner.ownedunits, &u)

	u.effectuser(true)

	return &u
}
