package io

type EventConf struct {
	Title    string
	Members  MembersConf
	Duties   DutiesConf
	Statuses Statuses
}

type MembersConf []Member

type Member struct {
	Name  string
	Count int
}

type DutiesConf []DayConf

type DayConf struct {
	Name  string
	Child []SubConf
}

type SubConf struct {
	Name  string
	Child []DutyConf
}

type DutyConf struct {
	Name    string
	Requier int
}

type Statuses [][]string

func ReadChouseisan(chouseiCSV string) (string, string, string) {
	event := ReadChouseisanCSV(chouseiCSV)

	tmpConf := getTempConf(event)

	memberConfStr := MarshalMembers(tmpConf.Members)
	dutiesConfStr := MarshalDuties(tmpConf.Duties)
	statusesStr := MarshalSutatuses(tmpConf.Statuses)

	return memberConfStr, dutiesConfStr, statusesStr
}

func getTempConf(event *Event) *EventConf {
	eventConf := &EventConf{
		Title: event.Title,
		Members: getTmpMemberConf(event.Members),
		Duties: getTmpDutiesConf(event.Schedule),
		Statuses: event.Statuses,
	}

	return eventConf
}

func getTmpMemberConf(members []string) MembersConf {
	membersConf := make(MembersConf, len(members))

	for i := range membersConf {
		membersConf[i] = Member{
			Name: members[i],
			Count: 0,
		}
	}

	return membersConf
}

func getTmpDutiesConf(schedule []string) DutiesConf {
	subConfs := make([]SubConf, len(schedule))
	for i := range subConfs {
		subConfs[i] = SubConf{
			Name: schedule[i],
			Child: []DutyConf{
				{
					Name: schedule[i],
					Requier: 1,
				},
			},
		}
	}

	DutiesConf := []DayConf{
		{
			Name: "day0",
			Child: subConfs,
		},
	}

	return DutiesConf
}
