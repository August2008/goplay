package graph

import (
)

type Vertex struct {
	Name string    
	Value interface{}
	Indegree, Outdegree []Edge
}

type Edge struct {
	u, v Vertex
	weight int
}

type Dag struct {
	vertices map[string]*Vertex
}

func NewDag() *Dag { 
	var d = new(Dag)
	return  d
}

func (d *Dag) AddVertex(x Vertex)  {
	if len(x.Name) > 0 {
		
	}
}

func (d *Dag) AddEdge(from, to string, weight int) {
	
}


