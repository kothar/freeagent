package freeagent

import (
	"log"
	"time"
)

type Timer struct {
	Running   bool   `json:"running"`
	StartFrom string `json:"start_from"`
}

type Timeslip struct {
	URL             string `json:"url,omitempty"`
	Task            string `json:"task"`
	User            string `json:"user"`
	Project         string `json:"project"`
	DatedOn         string `json:"dated_on"`
	Hours           string `json:"hours"`
	Comment         string `json:"comment"`
	BilledOnInvoice string `json:"billed_on_invoice"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Timer           *Timer `json:"timer"`
}

type timeslipDTO struct {
	Timeslip *Timeslip `json:"timeslip"`
}

func (c *FreeAgent) PostTimeslip(timeslip *Timeslip) (*Timeslip, error) {
	request := &timeslipDTO{timeslip}
	response := &timeslipDTO{}
	err := c.post("/timeslips", request, response)
	if err != nil {
		return nil, err
	}

	return response.Timeslip, nil
}

func (c *FreeAgent) GetTimeslip(id string) (*Timeslip, error) {
	result := &timeslipDTO{}
	err := c.get("/timeslips/"+id, result)
	if err != nil {
		return nil, err
	}

	return result.Timeslip, nil
}

type TimeslipView string

const (
	TimeslipViewAll      TimeslipView = "all"
	TimeslipViewUnbilled TimeslipView = "unbilled"
	TimeslipViewRunning  TimeslipView = "running"
)

type TimeslipQuery struct {
	FromDate     time.Time
	ToDate       time.Time
	UpdatedSince time.Time
	View         TimeslipView
}

func (c *FreeAgent) GetTimeslips(q *TimeslipQuery) ([]*Timeslip, error) {
	log.Fatal("Not implemented")
	return nil, nil
}
