package main

import "time"

// AkumuliDevops8Hosts produces Akumuli-specific queries for the devops groupby case.
type AkumuliDevops8Hosts struct {
	AkumuliDevops
}

func NewAkumuliDevops8Hosts(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newAkumuliDevopsCommon(dbConfig, start, end).(*AkumuliDevops)
	return &AkumuliDevops8Hosts{
		AkumuliDevops: *underlying,
	}
}

func (d *AkumuliDevops8Hosts) Dispatch(_, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageHourByMinuteEightHosts(q, scaleVar)
	return q
}
