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

	Khabib := fighters.NewFighter("Khabib Nurmagomedov", 32, 178)

	ufc.AddFighter(*JonJones, "Light-HeavyWeight")
	ufc.AddFighter(*Khabib, "LightWeight")

	mma.AddPromotion(ufc)
	pfl := &promotionsmanager.PFL{WeightClasses: make(map[string]*weightclasses.WeightClass)}

	Venom := fighters.NewFighter("Michael Venom", 28, 187)
	pfl.AddFighter(*Venom, "WelterWeight")

	mma.AddPromotion(pfl)

	for _, promotion := range mma.Promotions {
		promotion.NotifyFighters("New promotion announcement in " + promotion.Name())
	}

	//boxing

	boxing := combatsport.NewCombatSport("Boxing")
	goldenboy := &promotionsmanager.GoldenBoy{WeightClasses: make(map[string]*weightclasses.WeightClass)}

	Ali := fighters.NewFighter("Muhammad Ali", 25, 186)
	Frazier := fighters.NewFighter("George Frazier", 22, 185)

	goldenboy.AddFighter(*Ali, "HeavyWeight")
	goldenboy.AddFighter(*Frazier, "HeavyWeight")
	boxing.AddPromotion(goldenboy)

	for _, promotion := range boxing.Promotions {
		promotion.NotifyFighters("New promotion announcement in " + promotion.Name())
	}

}
