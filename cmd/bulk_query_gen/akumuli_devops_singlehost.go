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

// AkumuliDevopsSingleHostByHour produces Akumuli-specific queries for the devops single-host case.
type AkumuliDevopsSingleHostByHour struct {
	AkumuliDevops
}

func NewAkumuliDevopsSingleHostByHour(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newAkumuliDevopsCommon(dbConfig, start, end).(*AkumuliDevops)
	return &AkumuliDevopsSingleHostByHour{
		AkumuliDevops: *underlying,
	}
}

func (d *AkumuliDevopsSingleHostByHour) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageDayByHour(q, scaleVar)
	return q
}
