package weightclasses

import (
	fighters "fighter.com/event/Fighters"
)

type WeightClass struct {
	Name     string
	Fighters []fighters.Fighter
}

func NewWeightClass(name string) *WeightClass {
	return &WeightClass{Name: name, Fighters: make([]fighters.Fighter, 0)}
}

func (wc *WeightClass) AddFighter(fighter fighters.Fighter) {
	wc.Fighters = append(wc.Fighters, fighter)
}

func Shit() string {
	return "shite"
}
