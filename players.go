package main

// Player represents a player in the game.
type Player struct {
	Name                     string
	OwnedUnits               []*Unit
	UnitsLost                []*Unit //units lost since last turn
	FoodOutput               uint
	MetalOutput              uint
	ProductionOutput         uint
	ManagementOutput         uint
	FoodRequired             uint
	ManagementRequired       uint
	BuildQueue               []*UnitType
	BuildMetalRemaining      uint
	BuildProductionRemaining uint
	UnitsFinished            []*UnitType
	GameIn                   *Game
}

// ProcessUnitsLost processes the UnistLost list,
// adjusting "output" and "required" variables.
func (p *Player) ProcessUnitsLost() {
	for _, u := range p.UnitsLost {
		if u.Enabled {
			u.EffectUser(false)
		}
		u.Enabled = false
		p.GameIn.UnitMap[u.Location] = nil
		p.GameIn.MostRecentChanges[u.Location] = p.GameIn.CurrentTurnNum
		for i, u2 := range p.OwnedUnits {
			if u == u2 {
				p.OwnedUnits = append(p.OwnedUnits[:i], p.OwnedUnits[i+1:]...)
				break
			}
		}
	}
	p.UnitsLost = make([]*Unit, 0)
}

// UpdateActivations checks if the food and management requirements are less
// than supplied. if so, it deactivates buildings. if the food and management
// requirements are greater than supplied, it reactivates buildings.
func (p *Player) UpdateActivations() {
	if p.ManagementRequired < p.ManagementOutput {
		// find the unit that requires the most management (and isn't closed) and close it
		var maxmanagement uint
		var maxunit *Unit
		for _, u := range p.OwnedUnits {
			if u.Enabled && u.Stats.ManagementRequired > maxmanagement {
				maxunit = u
			}
		}
		maxunit.EffectUser(false)
		maxunit.Enabled = false
	} else if p.FoodRequired < p.FoodOutput {
		// find the unit that requires the most food (and isn't closed) and close it
		var maxfood uint
		var maxunit *Unit
		for _, u := range p.OwnedUnits {
			if u.Enabled && u.Stats.FoodRequired > maxfood {
				maxunit = u
			}
		}
		maxunit.EffectUser(false)
		maxunit.Enabled = false
	} else {
		// reactivate units
		for _, u := range p.OwnedUnits {
			if !u.Enabled && u.Stats.ManagementRequired+p.ManagementRequired < p.ManagementOutput &&
				u.Stats.FoodRequired+p.FoodRequired < p.FoodOutput {
				u.Enabled = true
				u.EffectUser(true)
			}
		}
	}
}

// UpdateBuilds updates the progress of a unit being built, and updates the
// build queue when applicable.
func (p *Player) UpdateBuilds() {
	if len(p.BuildQueue) > 0 {
		if p.BuildMetalRemaining > 0 {
			p.BuildMetalRemaining -= p.MetalOutput
		} else if p.BuildProductionRemaining > p.ProductionOutput {
			p.BuildProductionRemaining -= p.ProductionOutput
		} else {
			// build finished
			p.BuildProductionRemaining = 0
			p.UnitsFinished = append(p.UnitsFinished, p.BuildQueue[0])
			p.BuildQueue = p.BuildQueue[1:]
			if len(p.BuildQueue) > 0 {
				p.BuildMetalRemaining = p.BuildQueue[0].MetalCost
				p.BuildProductionRemaining = p.BuildQueue[0].ProductionCost
			}
		}
	}
}

// Turn is fired every turn.
func (p *Player) Turn() {
	p.ProcessUnitsLost()
	p.UpdateActivations()
	p.UpdateBuilds()
}

// placeunit places a unit from a player's finished units list.
// Returns whether successful or not.
func (p *Player) placeunit(u *UnitType, v Vec) bool {
	for i, unit := range p.UnitsFinished { // check if user can place that unit
		if unit == u {
			if NewUnit(u, v, p, p.GameIn) == nil {
				return false
			}
			p.UnitsFinished = append(p.UnitsFinished[:i], p.UnitsFinished[i+1:]...)
			return true
		}
	}
	return false
}

// MakePlayer constructs a new player with default values.
func MakePlayer(name string, game *Game) *Player {
	plr := Player{
		name,
		make([]*Unit, 0),
		make([]*Unit, 0),
		0,
		0,
		0,
		0,
		0,
		0,
		make([]*UnitType, 0),
		0,
		0,
		make([]*UnitType, 0),
		game,
	}
	game.PlayerList = append(game.PlayerList, plr)
	return &plr
}
