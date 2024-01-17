package mincostflow

func (g Graph) MinCostFlow(start, goal *Node, flows int) (int, any) {
	for ; flows > 0; flows-- {
		g.Dijkstra(start)
		if !goal.Used {
			return flows, goal.Pot
		}
		g.updatePotential()
		g.updateCapacity(start, goal)
	}

	return flows, goal.Pot
}

func (g Graph) updateCapacity(start, goal *Node) {
	tmp := goal

	for tmp.Prev != nil {
		tmp.Prev.Cap--
		tmp.Prev.Rev.Cap++
		tmp = tmp.Prev.From
	}
}

func (g Graph) updatePotential() {
	for _, node := range g {
		node.Pot = ci.AddCost(node.Pot, node.Dist)
	}
}
