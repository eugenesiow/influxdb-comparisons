package main

import "time"

// AkumuliDevopsSingleHost produces Akumuli-specific queries for the devops single-host case.
type AkumuliDevopsSingleHost struct {
	AkumuliDevops
}

func NewAkumuliDevopsSingleHost(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newAkumuliDevopsCommon(dbConfig, start, end).(*AkumuliDevops)
	return &AkumuliDevopsSingleHost{
		AkumuliDevops: *underlying,
	}
}

func (d *AkumuliDevopsSingleHost) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageHourByMinuteOneHost(q, scaleVar)
	return q
}
