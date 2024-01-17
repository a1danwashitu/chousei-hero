package io

import "gopkg.in/yaml.v2"

func MarshalMembers(members MembersConf) string {
	data, _ := yaml.Marshal(members)

	return string(data)
}

func MarshalDuties(duties DutiesConf) string {
	data, _ := yaml.Marshal(duties)

	return string(data)
}

func MarshalSutatuses(statuses Statuses) string {
	data, _ := yaml.Marshal(statuses)

	return string(data)
}

func UnmarshalMembers(membersStr string) MembersConf {
	membersConf := MembersConf{}
	_ = yaml.Unmarshal([]byte(membersStr), &membersConf)

	return membersConf
}

func UnmarshalDuties(dutiesStr string) DutiesConf {
	dutiesConf := DutiesConf{}
	_ = yaml.Unmarshal([]byte(dutiesStr), &dutiesConf)

	return dutiesConf
}

func UnmarshalStatuses(statusesStr string) Statuses {
	statusesConf := Statuses{}
	_ = yaml.Unmarshal([]byte(statusesStr), &statusesConf)

	return statusesConf
}
