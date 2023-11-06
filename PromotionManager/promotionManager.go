package promotionsmanager

import (
	fighters "fighter.com/event/Fighters"
	promotions "fighter.com/event/Promotions"
	weightclasses "fighter.com/event/WeightClasses"
)

type PromotionManager struct {
	Promotions map[string]promotions.Promotion
}

var instance *PromotionManager

func GetPromotionManager() *PromotionManager {
	if instance == nil {
		instance = &PromotionManager{
			Promotions: make(map[string]promotions.Promotion),
		}
	}
	return instance
}

// UFC, PFL, and IBF promotions implementing the Promotion interface
type UFC struct {
	WeightClasses map[string]*weightclasses.WeightClass
}

func (u UFC) Name() string {
	return "UFC"
}

// AddFighter adds a fighter to a weight class in UFC
func (ufc *UFC) AddFighter(fighter fighters.Fighter, weightClassName string) {
	weightClass, exists := ufc.WeightClasses[weightClassName]
	if !exists {
		weightClass = weightclasses.NewWeightClass(weightClassName)
		ufc.WeightClasses[weightClassName] = weightClass
	}
	weightClass.AddFighter(fighter)
}

// NotifyFighters notifies fighters in each weight class
func (ufc UFC) NotifyFighters(promotion string) {
	for _, weightClass := range ufc.WeightClasses {
		for _, fighter := range weightClass.Fighters {
			fighter.Update(promotion)
		}
	}
}

type PFL struct {
	WeightClasses map[string]*weightclasses.WeightClass
}

func (u PFL) Name() string {
	return "PFL"
}

// AddFighter adds a fighter to a weight class in UFC
func (ufc *PFL) AddFighter(fighter fighters.Fighter, weightClassName string) {
	weightClass, exists := ufc.WeightClasses[weightClassName]
	if !exists {
		weightClass = weightclasses.NewWeightClass(weightClassName)
		ufc.WeightClasses[weightClassName] = weightClass
	}
	weightClass.AddFighter(fighter)
}

// NotifyFighters notifies fighters in each weight class
func (ufc PFL) NotifyFighters(promotion string) {
	for _, weightClass := range ufc.WeightClasses {
		for _, fighter := range weightClass.Fighters {
			fighter.Update(promotion)
		}
	}
}

type GoldenBoy struct {
	WeightClasses map[string]*weightclasses.WeightClass
}

func (u GoldenBoy) Name() string {
	return "GoldenBoy"
}

// AddFighter adds a fighter to a weight class in UFC
func (ufc *GoldenBoy) AddFighter(fighter fighters.Fighter, weightClassName string) {
	weightClass, exists := ufc.WeightClasses[weightClassName]
	if !exists {
		weightClass = weightclasses.NewWeightClass(weightClassName)
		ufc.WeightClasses[weightClassName] = weightClass
	}
	weightClass.AddFighter(fighter)
}

// NotifyFighters notifies fighters in each weight class
func (ufc GoldenBoy) NotifyFighters(promotion string) {
	for _, weightClass := range ufc.WeightClasses {
		for _, fighter := range weightClass.Fighters {
			fighter.Update(promotion)
		}
	}
}
