package graph

type aStarNode struct {
	node   Node
	cost   float64
	rank   float64
	parent *aStarNode
	open   bool
	closed bool
	index  int
}

func (i aStarNode) Less(j QItem) bool {
	return i.rank < j.(aStarNode).rank
}

type nodeMap map[Node]*aStarNode

func (nm nodeMap) get(n Node) *aStarNode {
	a, ok := nm[n]
	if !ok {
		a = &aStarNode{
			node: n,
		}
		nm[n] = a
	}
	return a
}

func FindShortestPath(from, to Node) (path Path, distance float64, found bool) {
	nodeMap := nodeMap{}

	fromNode := nodeMap.get(from)

	queue := NewPriorityQueue()
	queue.Push(fromNode)

	for {
		if queue.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := queue.Pop().(*aStarNode)
		current.open = false
		current.closed = true

		if current == nodeMap.get(to) {
			// Found a path to the goal.
			p := Path{}
			curr := current
			for curr != nil {
				p = append(p, curr.node)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		for _, neighbor := range current.node.GetNeighbours() {
			cost := current.cost + current.node.GetCost(neighbor)
			neighborNode := nodeMap.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					queue.Pop()
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.GetHeuristicCost(to)
				neighborNode.parent = current
				queue.Push(neighborNode)
			}
		}
	}
}
