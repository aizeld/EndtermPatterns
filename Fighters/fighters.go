package fighters

import "fmt"

type FighterFactory struct{}

// CreateFighter creates a new fighter with the given attributes
func (ff FighterFactory) NewFighter(name string, age int, height int) *Fighter {
	return &Fighter{Name: name, Age: age, Height: height}
}

type Decorator interface {
	Modify(f *Fighter)
}

type ChampionStatusDecorator struct {
	IsChampion bool
}

func (csd ChampionStatusDecorator) Modify(f *Fighter) {
	f.Champion = csd.IsChampion
}

// Fighter represents a combatant
type Fighter struct {
	Name     string
	Age      int
	Height   int
	Champion bool
}

func (f *Fighter) Update(promotion string) {
	fmt.Printf("%s received promotion update: %s\n", f.Name, promotion)
}

func ShowChampions(fighters []*Fighter) {
	fmt.Println("Champion Fighters:")
	for _, fighter := range fighters {
		if fighter.Champion {
			fmt.Printf("Name: %s, Age: %d, Height: %d\n", fighter.Name, fighter.Age, fighter.Height)
		}
	}
}
