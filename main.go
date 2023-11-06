package main

import (
	combatsport "fighter.com/event/CombatSport"
	fighters "fighter.com/event/Fighters"
	promotionsmanager "fighter.com/event/PromotionManager"
	weightclasses "fighter.com/event/WeightClasses"
)

func main() {

	mma := combatsport.NewCombatSport("MMA")

	ufc := &promotionsmanager.UFC{WeightClasses: make(map[string]*weightclasses.WeightClass)}

	JonJones := fighters.NewFighter("Jon Jones", 20, 194)

	ufc.AddFighter(*JonJones, "Light-Heavyweight")

	mma.AddPromotion(ufc)
	for _, promotion := range mma.Promotions {
		promotion.NotifyFighters("New promotion announcement in " + promotion.Name())
	}
}
