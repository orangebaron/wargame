package core

// UnitType contains data for a type of unit.
// a unit can be a boat or a tile.
type UnitType struct {
	Name               string
	FoodOutput         uint
	MetalOutput        uint
	ProductionOutput   uint
	ManagementOutput   uint
	FoodRequired       uint
	ManagementRequired uint
	MetalCost          uint
	ProductionCost     uint
	MaxHealth          uint
	Speed              uint
	HitRange           uint
	Damage             uint
}

// Types is a list of all possible unit types.
var Types = [14]UnitType{
	UnitType{"Farm", 75, 0, 0, 0, 5, 1, 1, 20, 3, 0, 0, 0},
	UnitType{"Urban Farm", 200, 0, 0, 0, 10, 3, 5, 50, 1, 0, 0, 0},
	UnitType{"Mineshaft", 0, 1, 0, 0, 30, 7, 3, 15, 3, 0, 0, 0},
	UnitType{"Factory", 0, 0, 1, 0, 40, 10, 10, 60, 1, 0, 0, 0},
	UnitType{"Robotic Factory", 0, 0, 3, 0, 0, 7, 20, 120, 1, 0, 0, 0},
	UnitType{"Small Office", 0, 0, 0, 100, 20, 0, 4, 40, 1, 0, 0, 0},
	UnitType{"Office Tower", 0, 0, 0, 300, 50, 0, 10, 100, 1, 0, 0, 0},
	UnitType{"Wall", 0, 0, 0, 0, 0, 0, 10, 15, 10, 0, 0, 0},
	UnitType{"Turret", 0, 0, 0, 0, 2, 1, 15, 30, 5, 0, 1, 5},
	UnitType{"Missile Station", 0, 0, 0, 0, 5, 2, 20, 40, 5, 0, 20, 10},
	UnitType{"Cruiser", 0, 0, 0, 0, 10, 4, 40, 90, 5, 5, 5, 7},
	UnitType{"Frigate", 0, 0, 0, 0, 7, 3, 20, 50, 4, 7, 2, 4},
	UnitType{"Destroyer", 0, 0, 0, 0, 10, 4, 35, 90, 5, 10, 2, 5},
	UnitType{"Aircraft Carrier", 0, 0, 0, 0, 50, 6, 50, 120, 15, 5, 10, 10},
}

// Unit represents a unit or tile.
type Unit struct {
	Stats    *UnitType
	Health   uint
	Location Vec
	Owner    *Player
	Enabled  bool
	GameIn   *Game
}

func (u *Unit) move(v Vec) bool {
	if u.GameIn.UnitMap[v] == nil && u.Enabled {
		u.Location = v
		return true
	}
	return false
}

// EffectUser applies output and requirement effects to the player's values
func (u *Unit) EffectUser(positive bool) {
	p := u.Owner
	if positive {
		p.FoodOutput += u.Stats.FoodOutput
		p.MetalOutput += u.Stats.MetalOutput
		p.ProductionOutput += u.Stats.ProductionOutput
		p.ManagementOutput += u.Stats.ManagementOutput
		p.FoodRequired += u.Stats.FoodRequired
		p.ManagementRequired += u.Stats.ManagementRequired
	} else {
		p.FoodOutput -= u.Stats.FoodOutput
		p.MetalOutput -= u.Stats.MetalOutput
		p.ProductionOutput -= u.Stats.ProductionOutput
		p.ManagementOutput -= u.Stats.ManagementOutput
		p.FoodRequired -= u.Stats.FoodRequired
		p.ManagementRequired -= u.Stats.ManagementRequired
	}
}

// Damage deals damage to a unit, and deals with killing it if applicable.
func (u *Unit) Damage(d uint) bool { //return value: whether unit was killed
	if u.Health <= d {
		u.Health = 0
		u.Owner.UnitsLost = append(u.Owner.UnitsLost, u) // register with owner that unit is dead
		return true
	}
	u.Health -= d
	return false
}

// Attack causes unit u to attack target.
func (u *Unit) Attack(target *Unit) bool {
	if !u.Enabled {
		return false
	}
	return target.Damage(u.Stats.Damage)
}

// NewUnit creates and initializes a new unit with the given specifications.
func NewUnit(stats *UnitType, location Vec, owner *Player, game *Game) *Unit {
	if game.UnitMap[location] != nil {
		return nil
	}

	u := Unit{stats, stats.MaxHealth, location, owner, true, game}
	game.UnitMap[location] = &u
	owner.OwnedUnits = append(owner.OwnedUnits, &u)

	u.EffectUser(true)

	return &u
}
