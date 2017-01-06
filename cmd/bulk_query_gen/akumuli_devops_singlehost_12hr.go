package main

import "time"

// AkumuliDevopsSingleHost12hr produces Akumuli-specific queries for the devops single-host case over a 12hr period.
type AkumuliDevopsSingleHost12hr struct {
	AkumuliDevops
}

func NewAkumuliDevopsSingleHost12hr(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newAkumuliDevopsCommon(dbConfig, start, end).(*AkumuliDevops)
	return &AkumuliDevopsSingleHost12hr{
		AkumuliDevops: *underlying,
	}
}

func (d *AkumuliDevopsSingleHost12hr) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsage12HoursByMinuteOneHost(q, scaleVar)
	return q
}
