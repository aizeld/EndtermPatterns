package promotions

import (
	fighters "fighter.com/event/Fighters"
)

type Promotion interface {
	Name() string
	AddFighter(fighter fighters.Fighter, weightClassName string)
	NotifyFighters(promotion string)
}
