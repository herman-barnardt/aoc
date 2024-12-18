package graph

type Node interface {
	GetNeighbours() []Node
	GetCost(to Node) float64
	GetHeuristicCost(to Node) float64
	Equal(to Node) bool
}

type Path []Node

type BasicNode struct {
	X, Y               int
	Value              string
	PossibleNeighbours []*BasicNode
	NeighbourFilter    func(*BasicNode) bool
}

func (p *BasicNode) GetNeighbours() []Node {
	nodeNeighbours := make([]Node, 0)
	for _, neighbour := range p.PossibleNeighbours {
		if p.NeighbourFilter(neighbour) {
			nodeNeighbours = append(nodeNeighbours, neighbour)
		}
	}
	return nodeNeighbours
}
func (p *BasicNode) GetCost(to Node) float64 {
	return 1
}

func (p *BasicNode) GetHeuristicCost(to Node) float64 {
	return 1
}

func (p *BasicNode) Equal(to Node) bool {
	return p.X == to.(*BasicNode).X && p.Y == to.(*BasicNode).Y
}
