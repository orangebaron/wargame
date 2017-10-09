package main

type time uint    //HOW TO REPRESENT TIME?
func now() time { //HOW TO GET NOW'S TIME?
	return 0
}

// unitType contains data for a type of unit or tile.
type unitType struct {
	name        string
	foodO       uint
	metalO      uint
	productionO uint
	managementO uint
	peopleR     uint
	managementR uint
	metalC      uint
	productionC uint
	maxhealth   uint
	speed       uint
	hitrange    uint
	damage      uint
}

// types is a list of all possible types.
var types = [14]unitType{
	unitType{"Farm", 75, 0, 0, 0, 5, 1, 1, 20, 3, 0, 0, 0},
	unitType{"Urban Farm", 200, 0, 0, 0, 10, 3, 5, 50, 1, 0, 0, 0},
	unitType{"Mineshaft", 0, 1, 0, 0, 30, 7, 3, 15, 3, 0, 0, 0},
	unitType{"Factory", 0, 0, 1, 0, 40, 10, 10, 60, 1, 0, 0, 0},
	unitType{"Robotic Factory", 0, 0, 3, 0, 0, 7, 20, 120, 1, 0, 0, 0},
	unitType{"Small Office", 0, 0, 0, 100, 20, 0, 4, 40, 1, 0, 0, 0},
	unitType{"Office Tower", 0, 0, 0, 300, 50, 0, 10, 100, 1, 0, 0, 0},
	unitType{"Wall", 0, 0, 0, 0, 0, 0, 10, 15, 10, 0, 0, 0},
	unitType{"Turret", 0, 0, 0, 0, 2, 1, 15, 30, 5, 0, 1, 5},
	unitType{"Missile Station", 0, 0, 0, 0, 5, 2, 20, 40, 5, 0, 20, 10},
	unitType{"Cruiser", 0, 0, 0, 0, 10, 4, 40, 90, 5, 5, 5, 7},
	unitType{"Frigate", 0, 0, 0, 0, 7, 3, 20, 50, 4, 7, 2, 4},
	unitType{"Destroyer", 0, 0, 0, 0, 10, 4, 35, 90, 5, 10, 2, 5},
	unitType{"Aircraft Carrier", 0, 0, 0, 0, 50, 6, 50, 120, 15, 5, 10, 10},
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
	utype    *unitType
	health   uint
	location vec
}

// unitmap is a map of a location on the board to the unit in that location.
var unitmap = make(map[vec]*unit)

// mostrecentchanges list the most recent time a spot has been changed.
var mostrecentchanges = make(map[vec]time)

// killevents list events to be called when a unit is killed.
var killevents = make(map[*unit][]func(*unit))

func (u *unit) move(v vec) bool {
	if unitmap[v] == nil {
		u.location = v
		return true
	}
	return false
}
func (u *unit) registerkillevent(f func(*unit)) {
	killevents[u] = append(killevents[u], f)
}
func (u *unit) kill() {
	u.health = 0
	for _, event := range killevents[u] {
		event(u)
	}
}
func (u *unit) damage(d uint) bool { //return value: whether unit was killed
	if u.health <= d {
		u.kill()
		return true
	}
	u.health -= d
	return false
}
func (u *unit) attack(target *unit) bool {
	return target.damage(u.utype.damage)
}

func main() {

}
