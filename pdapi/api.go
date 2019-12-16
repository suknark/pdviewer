package pdapi

import (
	"github.com/PagerDuty/go-pagerduty"
	"log"
	"os"
	"strings"
	"pdviewer/tools"
)

var authtoken = os.Getenv("PDTOKEN")
var schedules = os.Getenv("SHEDULES")


func GetIncidents() ([]tools.Inc, bool) {
	t := true
	var ii []tools.Inc
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

	client := pagerduty.NewClient(authtoken)
	if eps, err := client.ListIncidents(opts); err != nil {
		log.Println(err)
	} else {
		for _, p := range eps.Incidents {
			if p.Status != "resolved" {
				if p.Urgency == "high" {
					t = false
				}
				i := tools.Inc{
					Urg:     p.Urgency,
					Summary: strings.Replace(p.APIObject.Summary, "\n", "", -1),
					Time:    strings.Replace(p.Assignments[0].Assignee.Summary, "\n", "", -1),
				}
				ii = append(ii, i)
			}
		}
	}
	return ii, t
}

func OnCall() string {
	opts := pagerduty.ListOnCallOptions{}
	client := pagerduty.NewClient(authtoken)
	if eps, err := client.ListOnCalls(opts); err != nil {
		log.Println(err)
	} else {
		for _, p := range eps.OnCalls {
			if p.Schedule.Summary == schedules {
				return p.User.Summary
			} else {
				continue
			}
		}
	}
	return "oops"
}
