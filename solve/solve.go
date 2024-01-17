package solve

import (
	"github.com/a1danwashitu/chousei-hero/io"
	"github.com/a1danwashitu/chousei-hero/mincostflow"
)

func Solve(event *io.EventConf) []*io.Assignment {
	g := buildGraph(event)

	mg := buildMinCostFlowGraph(g)
	mincostflow.Init(costInterface{})
	mg.MinCostFlow(mg[0], mg[1], g.getRequireFlows())

	assignments := g.getAssignments()

	return assignments
}

func buildMinCostFlowGraph(g *graph) mincostflow.Graph {
	mg := mincostflow.Graph{}

	mg = append(mg, g.start)
	mg = append(mg, g.goal)
	for _, member := range g.members {
		mg = append(mg, member.node)
		for _, day := range member.days {
			mg = append(mg, day.node)
			for _, sub := range day.subs {
				mg = append(mg, sub.node)
			}
		}
	}
	for _, duty := range g.duties {
		mg = append(mg, duty.node)
	}

	return mg
}

func (g *graph) getRequireFlows() int {
	flows := 0
	for _, duty := range g.duties {
		flows += duty.require
	}

	return flows
}
