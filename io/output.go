package io

import (
	"strings"
)

type Assignment struct {
	Duty      string
	Assignees []string
	Require   int
}

func OutputResult(assignments []*Assignment) string {
	res := ""
	for _, assign := range assignments {
		assignees := strings.Join(assign.Assignees, ", ")
		row := assign.Duty + ": " + assignees + "\n"
		res += row
	}

	return res
}
