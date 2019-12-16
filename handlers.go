package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type Handlers struct {
	template    *template.Template
	uptimeSince time.Time
	pd          *PdApi
}

func NewHandlers(pd *PdApi) *Handlers {
	templateData, err := ioutil.ReadFile("templates/index.html")
	if err != nil {
		panic(fmt.Errorf("cannot read tempalte: %s", err))
	}
	h := &Handlers{
		template:    template.Must(template.New("index").Parse(string(templateData))),
		uptimeSince: time.Now(),
		pd:          pd,
	}
	return h
}

func (h *Handlers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	incidents, t := h.pd.GetIncidents()
	if !t {
		h.uptimeSince = time.Now()
	}
	oncall := h.pd.OnCall()
	uptime := time.Since(h.uptimeSince)
	data := struct {
		Incidents []Incident
		OnCall    string
		Uptime    time.Duration
		T         bool
	}{incidents, oncall, uptime, t}

	if err := h.template.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
