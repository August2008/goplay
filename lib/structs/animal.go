package structs

//AnimalType type
type AnimalType string

const (
	mamals   AnimalType = "Mamal"
	reptiles AnimalType = "Reptile"
	birds    AnimalType = "Bird"
	insect   AnimalType = "Insect"
	aquatic  AnimalType = "Aquatic"
)

//Animal struct
type Animal struct {
	animalType AnimalType
}

//Type func
func (obj *Animal) Type() (t AnimalType) {
	return obj.animalType
}

//Taxonomy function
func (obj *Animal) Taxonomy() string {
	return string(obj.animalType)
}
