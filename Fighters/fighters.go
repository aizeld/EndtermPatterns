package fighters

import "fmt"

// Fighter represents a combatant
type Fighter struct {
	Name   string
	Age    int
	Height int
}

func NewFighter(name string, age int, height int) *Fighter {
	return &Fighter{Name: name, Age: age, Height: height}
}

func (f *Fighter) Update(promotion string) {
	fmt.Printf("%s received promotion update: %s\n", f.Name, promotion)
}
