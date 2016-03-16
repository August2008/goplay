package structs

import "fmt"

//OwlType type
type OwlType string

const (
	snowyOwl OwlType = "Snowy Owl"
	eagleOwl OwlType = "Eagle Owl"
)

//Owl struct
type Owl struct {
	Bird
	owlType OwlType
}

//Type function
func (ot *Owl) Type() (result OwlType) {
	return ot.owlType
}

//NewOwl constructor
func NewOwl() (obj Owl) {
	obj = Owl{
		Bird{
			Animal{"Bird"},
			"Owl",
		},
		"Snowy Owl",
	}
	return
}

//Taxonomy function
func (obj *Owl) Taxonomy() string {
	return fmt.Sprint(obj.animalType, "->", obj.birdType, "->", obj.owlType)
}
