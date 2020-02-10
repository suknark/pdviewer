package main

import (
	"log"
	"strings"

	"github.com/PagerDuty/go-pagerduty"
)

type Incident struct {
	Urg     string
	Summary string
	Time    string
}

type PdApi struct {
	schedules string
	client    *pagerduty.Client
}

func NewPdApi(token, schedules string) *PdApi {
	client := pagerduty.NewClient(token)
	return &PdApi{
		client:    client,
		schedules: schedules,
	}
}

func (pd *PdApi) GetIncidents() ([]Incident, bool) {
	t := true
	var ii []Incident
	statuses := []string{"triggered", "acknowledged"}
	nolimit := pagerduty.APIListObject{
		Limit: 100000,
		More:  true,
	}
	opts := pagerduty.ListIncidentsOptions{
		APIListObject: nolimit,
		SortBy:        "incident_number:desc",
		Statuses:      statuses,
	}
	eps, err := pd.client.ListIncidents(opts)
	if err != nil {
		log.Println(err)
		return nil, t
	}
	for _, p := range eps.Incidents {
		if p.Status != "resolved" && p.EscalationPolicy.Summary == pd.schedules {
			if p.Urgency == "high" {
				t = false
			}
			i := Incident{
				Urg:     p.Urgency,
				Summary: strings.Replace(p.APIObject.Summary, "\n", "", -1),
				Time:    strings.Replace(p.Assignments[0].Assignee.Summary, "\n", "", -1),
			}
			ii = append(ii, i)
		}
	}
	return ii, t
}

func (pd *PdApi) OnCall() string {
	opts := pagerduty.ListOnCallOptions{}
	eps, err := pd.client.ListOnCalls(opts)
	if err != nil {
		log.Println(err)
		return "oops"
	}
	for _, p := range eps.OnCalls {
		if p.Schedule.Summary == pd.schedules {
			return p.User.Summary
		} else {
			continue
		}
	}
	return "oops"
}
