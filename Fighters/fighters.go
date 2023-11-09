package fighters

import "fmt"

type FighterFactory struct{}

// CreateFighter creates a new fighter with the given attributes
func (ff FighterFactory) NewFighter(name string, age int, height int) *Fighter {
	return &Fighter{Name: name, Age: age, Height: height}
}

// Fighter represents a combatant
type Fighter struct {
	Name   string
	Age    int
	Height int
}

func (f *Fighter) Update(promotion string) {
	fmt.Printf("%s received promotion update: %s\n", f.Name, promotion)
}
