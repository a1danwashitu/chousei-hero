package solve

import (
	"github.com/a1danwashitu/chousei-hero/io"
	"github.com/a1danwashitu/chousei-hero/mincostflow"
)

const typeStart string = "start"
const typeGoal string = "goal"
const typeMember string = "member"
const typeDay string = "day"
const typeSub string = "sub"
const typeDuty string = "duty"

type graph struct {
	start   *mincostflow.Node
	goal    *mincostflow.Node
	members []*member
	duties  []*duty
}

type member struct {
	node  *mincostflow.Node
	id    int
	name  string
	count int
	days  []*day
}

type day struct {
	node *mincostflow.Node
	id   int
	name string
	subs []*sub
}

type sub struct {
	node   *mincostflow.Node
	id     int
	name   string
	duties []*duty
}

type duty struct {
	node     *mincostflow.Node
	id       int
	name     string
	require int
}

func buildGraph(event *io.EventConf) *graph {
	g := &graph{}
	g.buildNode(event)
	g.buildEdge(event)

	return g
}

func (g *graph) buildNode(event *io.EventConf) {
	g.start = &mincostflow.Node{
		Type:  "start",
		Id:    0,
		Nexts: []*mincostflow.Edge{},
	}

	g.goal = &mincostflow.Node{
		Type:  "goal",
		Id:    0,
		Nexts: []*mincostflow.Edge{},
	}

	g.members = buildNodeMembers(event)

	g.duties = buildNodeDuties(event.Duties, len(event.Statuses))

	for _, member := range g.members {
		id := 0
		for _, day := range member.days {
			for _, sub := range day.subs {
				for _, subDuty := range sub.duties {
					subDuty.node = g.duties[id].node
					id++
				}
			}
		}
	}
}

func buildNodeMembers(eventConf *io.EventConf) []*member {
	members := make([]*member, len(eventConf.Members))
	for i := range members {
		members[i] = &member{
			node: &mincostflow.Node{
				Type:  typeMember,
				Id:    i,
				Nexts: []*mincostflow.Edge{},
			},
			id:    i,
			name:  eventConf.Members[i].Name,
			count: eventConf.Members[i].Count,
			days:  buildNodeDays(eventConf.Duties),
		}
	}

	return members
}

func buildNodeDays(dutiesConf io.DutiesConf) []*day {
	days := make([]*day, len(dutiesConf))
	for i := range days {
		days[i] = &day{
			node: &mincostflow.Node{
				Type:  typeDay,
				Id:    i,
				Nexts: []*mincostflow.Edge{},
			},
			id:   i,
			name: dutiesConf[i].Name,
			subs: buildNodeSubs(dutiesConf[i].Child),
		}
	}

	return days
}

func buildNodeSubs(subConfs []io.SubConf) []*sub {
	subs := make([]*sub, len(subConfs))
	for i := range subs {
		subs[i] = &sub{
			node: &mincostflow.Node{
				Type:  typeSub,
				Id:    i,
				Nexts: []*mincostflow.Edge{},
			},
			id:     i,
			name:   subConfs[i].Name,
			duties: buildNodeSubDuties(subConfs[i].Child),
		}
	}

	return subs
}

func buildNodeSubDuties(dutyConfs []io.DutyConf) []*duty {
	duties := make([]*duty, len(dutyConfs))
	for i := range duties {
		duties[i] = &duty{
			id:       i,
			name:     dutyConfs[i].Name,
			require: dutyConfs[i].Requier,
		}
	}

	return duties
}

func buildNodeDuties(dutiesConf io.DutiesConf, dutyN int) []*duty {
	duties := make([]*duty, dutyN)
	id := 0
	for _, day := range dutiesConf {
		for _, sub := range day.Child {
			for _, subDuty := range sub.Child {
				duties[id] = &duty{
					node: &mincostflow.Node{
						Type:  typeDuty,
						Id:    id,
						Nexts: []*mincostflow.Edge{},
					},
					id:       id,
					name:     subDuty.Name,
					require: subDuty.Requier,
				}
				id++
			}
		}
	}

	return duties
}

func (g *graph) buildEdge(event *io.EventConf) {
	g.buildEdgeStartToMember()
	g.buildEdgeMemberToDay()
	g.buildEdgeDayToSub()
	g.buildEdgeSubToDuty(event.Statuses)
	g.buildEdgeDutyToGoal()
}

func (g *graph) buildEdgeStartToMember() {
	for _, member := range g.members {
		for i := 0; i < len(g.duties); i++ {
			pos, neg := getEdgePair(g.start, member.node, 1, 0, i+member.count, i, 0)
			g.start.Nexts = append(g.start.Nexts, pos)
			member.node.Nexts = append(member.node.Nexts, neg)
		}
	}
}

func (g *graph) buildEdgeMemberToDay() {
	for _, member := range g.members {
		for _, day := range member.days {
			for i := 0; i < len(day.subs); i++ {
				pos, neg := getEdgePair(member.node, day.node, 1, i, 0, 0, 0)
				member.node.Nexts = append(member.node.Nexts, pos)
				day.node.Nexts = append(day.node.Nexts, neg)
			}
		}
	}
}

func (g *graph) buildEdgeDayToSub() {
	for _, member := range g.members {
		for _, day := range member.days {
			for _, sub := range day.subs {
				pos, neg := getEdgePair(day.node, sub.node, 1, 0, 0, 0, 0)
				day.node.Nexts = append(day.node.Nexts, pos)
				sub.node.Nexts = append(sub.node.Nexts, neg)
			}
		}
	}
}

func (g *graph) buildEdgeSubToDuty(statuses io.Statuses) {
	for mid, member := range g.members {
		for _, day := range member.days {
			for _, sub := range day.subs {
				for _, duty := range sub.duties {
					did := duty.node.Id
					level := statusToLevel(statuses[did][mid])
					if level < 0 {
						continue
					}
					pos, neg := getEdgePair(sub.node, duty.node, 1, 0, 0, 0, level)
					sub.node.Nexts = append(sub.node.Nexts, pos)
					duty.node.Nexts = append(duty.node.Nexts, neg)
				}
			}
		}
	}
}

func (g *graph) buildEdgeDutyToGoal() {
	for _, duty := range g.duties {
		pos, neg := getEdgePair(duty.node, g.goal, duty.require, 0,0,0,0)
		duty.node.Nexts = append(duty.node.Nexts, pos)
		g.goal.Nexts = append(g.goal.Nexts, neg)
	}
}

func getEdgePair(from, to *mincostflow.Node, cap, day, total, tmp, level int) (*mincostflow.Edge, *mincostflow.Edge) {
	costPos, costNeg := getCostPair(day, total, tmp, level)

	positive := getEdge(from, to, cap, costPos)
	negative := getEdge(to, from, 0, costNeg)

	positive.Rev = negative
	negative.Rev = positive

	return positive, negative
}

func getEdge(from, to *mincostflow.Node, cap int, cos *cost) *mincostflow.Edge {
	return &mincostflow.Edge{
		Cost: cos,
		Cap:  cap,
		From: from,
		To:   to,
	}
}

func getCostPair(day, total, tmp, level int) (*cost, *cost) {
	return &cost{
			Day:   day,
			Total: total,
			Tmp:   tmp,
			Level: level,
		}, &cost{
			Day:   -day,
			Total: -total,
			Tmp:   -tmp,
			Level: -level,
		}
}

func statusToLevel(status string) int {
	switch status {
	case "◯":
		return 0
	case "△":
		return 1
	case "×":
		return -1
	}
	return -1
}
