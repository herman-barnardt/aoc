package graph

type Node interface {
	GetNeighbours() []Node
	GetCost(to Node) float64
	GetHeuristicCost(to Node) float64
	Equal(to Node) bool
}

type Path []Node
