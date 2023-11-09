package main

import (
	combatsport "fighter.com/event/CombatSport"
	fighters "fighter.com/event/Fighters"
	promotionsmanager "fighter.com/event/PromotionManager"
	weightclasses "fighter.com/event/WeightClasses"
)

func main() {
	fighterfactory := fighters.FighterFactory{}
	mma := combatsport.NewCombatSport("MMA")

	ufc := &promotionsmanager.UFC{WeightClasses: make(map[string]*weightclasses.WeightClass)}
	promotionsmanager.GetPromotionManager().Promotions["UFC"] = ufc

	JonJones := fighterfactory.NewFighter("Jon Jones", 20, 194)

	Khabib := fighterfactory.NewFighter("Khabib Nurmagomedov", 32, 178)

	ufc.AddFighter(*JonJones, "Light-HeavyWeight")
	ufc.AddFighter(*Khabib, "LightWeight")

	mma.AddPromotion(ufc)
	pfl := &promotionsmanager.PFL{WeightClasses: make(map[string]*weightclasses.WeightClass)}
	promotionsmanager.GetPromotionManager().Promotions["PFL"] = pfl

	Venom := fighterfactory.NewFighter("Michael Venom", 28, 187)
	pfl.AddFighter(*Venom, "WelterWeight")

	mma.AddPromotion(pfl)

	for _, promotion := range mma.Promotions {
		promotion.NotifyFighters("New promotion announcement in " + promotion.Name())
	}

	championDecorator := fighters.ChampionStatusDecorator{IsChampion: true}

	//boxing

	boxing := combatsport.NewCombatSport("Boxing")
	goldenboy := &promotionsmanager.GoldenBoy{WeightClasses: make(map[string]*weightclasses.WeightClass)}
	promotionsmanager.GetPromotionManager().Promotions["GoldenBoy"] = goldenboy

	Ali := fighterfactory.NewFighter("Muhammad Ali", 25, 186)
	Frazier := fighterfactory.NewFighter("George Frazier", 22, 185)

	goldenboy.AddFighter(*Ali, "HeavyWeight")
	goldenboy.AddFighter(*Frazier, "HeavyWeight")
	boxing.AddPromotion(goldenboy)

	championDecorator.Modify(Khabib)
	championDecorator.Modify(JonJones)
	championDecorator.Modify(Ali)

	fighterss := []*fighters.Fighter{Khabib, JonJones, Ali, Frazier}

	fighters.ShowChampions(fighterss)

	for _, promotion := range boxing.Promotions {
		promotion.NotifyFighters("New promotion announcement in " + promotion.Name())
	}

}
