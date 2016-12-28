package algorithms

type Planet struct {
	Name string
}

type Planets []Planet

var planets = Planets{
	{Name: "Mars"},
	{Name: "Jupiter"},
	{Name: "Saturn"},
	{Name: "Pluto"},
	{Name: "Earth"},
	{Name: "Venus"},
	{Name: "Uranus"},
	{Name: "Neptune"},
}

type Vertex struct {
	Name   string
	Weight float64
}

type Edge struct {
	From   string
	To     []string
	Weight float64
}
