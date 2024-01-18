package io

import (
	"bufio"
	"encoding/csv"
	"strings"
)

type Event struct {
	Title    string
	Schedule []string
	Members  []string
	Statuses [][]string
	Coments  []string
}

func ReadChouseisanCSV(chouseiCSV string) *Event {
	event := &Event{}

	r := csv.NewReader(bufio.NewReader(strings.NewReader(chouseiCSV)))
	r.FieldsPerRecord = -1

	row, _ := r.Read()
	event.Title = row[0]

	table, _ := r.ReadAll()

	event.Schedule = make([]string, len(table)-2)
	for i := range event.Schedule {
		event.Schedule[i] = table[i+1][0]
	}

	event.Members = table[0][1:]

	event.Statuses = make([][]string, len(event.Schedule))
	for i := range event.Statuses {
		event.Statuses[i] = table[i+1][1:]
	}

	event.Coments = table[len(event.Schedule)-1][1:]

	return event
}
