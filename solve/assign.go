package solve

import (
	"github.com/a1danwashitu/chousei-hero/io"
)

func (g *graph) getAssignments() []*io.Assignment {
	assignments := make([]*io.Assignment, len(g.duties))
	for i := range assignments {
		duty := g.duties[i]

		assignments[i] = &io.Assignment{
			Duty:      duty.name,
			Assignees: make([]string, 0, duty.require),
			Require:   duty.require,
		}
	}

	for _, member := range g.members {
		memberId := member.id
		for _, edgeMtoD := range member.node.Nexts {
			if edgeMtoD.To.Type == typeDay && edgeMtoD.Cap == 0 {
				dayNode := edgeMtoD.To
				for _, edgeDtoS := range dayNode.Nexts {
					if edgeDtoS.To.Type == typeSub && edgeDtoS.Cap == 0 {
						subNode := edgeDtoS.To
						for _, edgeStoD := range subNode.Nexts {
							if edgeStoD.To.Type == typeDuty && edgeStoD.Cap == 0 {
								dutyId := edgeStoD.To.Id
								assignments[dutyId].Assignees = append(assignments[dutyId].Assignees, g.members[memberId].name)
							}
						}
					}
				}
			}
		}
	}

	return assignments
}
