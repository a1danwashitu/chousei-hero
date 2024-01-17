package mincostflow

import (
	"container/heap"
)

func (g Graph) Dijkstra(start *Node) {
	g.initDist()
	g.initUsed()
	start.Dist = ci.NewCost()
	h := &Heap{
		{To: start, Cost: ci.NewCost(), Cap: 1},
	}
	heap.Init(h)

	for h.Len() > 0 {
		edge := heap.Pop(h).(*Edge)
		tmp := edge.To
		if tmp.Used {
			continue
		}
		tmp.Used = true
		for _, next := range tmp.Nexts {
			if next.Cap <= 0 {
				continue
			}
			cost := ci.SubCost(ci.AddCost(ci.AddCost(tmp.Dist, next.Cost), tmp.Pot), next.To.Pot)
			if ci.Less(cost, next.To.Dist) {
				next.To.Dist = cost
				next.To.Prev = next
				heap.Push(h, &Edge{To: next.To, Cost: cost, Cap: 1})
			}
		}
	}
}

func (g Graph) initDist() {
	for _, node := range g {
		node.Dist = ci.NewMaxCost()
	}
}

func (g Graph) initUsed() {
	for _, node := range g {
		node.Used = false
	}
}

type Heap []*Edge

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return ci.Less(h[i].Cost, h[j].Cost) }

func (h *Heap) Push(e interface{}) {
	*h = append(*h, e.(*Edge))
}

func (h *Heap) Pop() interface{} {
	old := *h
	length := len(old)
	elm := old[length-1]
	*h = old[:length-1]
	return elm
}
