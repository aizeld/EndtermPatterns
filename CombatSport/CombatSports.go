package combatsport

import (
	promotions "fighter.com/event/Promotions"
)

type CombatSport struct {
	Name       string
	Promotions []promotions.Promotion
}

func NewCombatSport(name string) *CombatSport {
	return &CombatSport{Name: name, Promotions: make([]promotions.Promotion, 0)}
}

func (cs *CombatSport) AddPromotion(promotion promotions.Promotion) {
	cs.Promotions = append(cs.Promotions, promotion)
}
