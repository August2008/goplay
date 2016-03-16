package structs

import "fmt"

//BirdType type
type BirdType string

const (
	owl   BirdType = "Owl"
	eagle BirdType = "Eagle"
	duck  BirdType = "Duck"
)

//Bird struct
type Bird struct {
	Animal
	birdType BirdType
}

//Type function
func (bt *Bird) Type() (t BirdType) {
	return bt.birdType
}

//Taxonomy function
func (bt *Bird) Taxonomy() string {
	return fmt.Sprint(bt.animalType, "->", bt.birdType)
}
