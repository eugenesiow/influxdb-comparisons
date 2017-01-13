package main

import "time"

// InfluxDevopsSingleHost produces Influx-specific queries for the devops single-host case.
type InfluxDevopsSingleHost struct {
	InfluxDevops
}

func NewInfluxDevopsSingleHost(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newInfluxDevopsCommon(dbConfig, start, end).(*InfluxDevops)
	return &InfluxDevopsSingleHost{
		InfluxDevops: *underlying,
	}
}

func (d *InfluxDevopsSingleHost) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageHourByMinuteOneHost(q, scaleVar)
	return q
}

type InfluxDevopsSingleHostByHour struct {
	InfluxDevops
}

func NewInfluxDevopsSingleHostByHour(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	underlying := newInfluxDevopsCommon(dbConfig, start, end).(*InfluxDevops)
	return &InfluxDevopsSingleHostByHour{
		InfluxDevops: *underlying,
	}
}

func (d *InfluxDevopsSingleHostByHour) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	d.MaxCPUUsageDayByHour(q, scaleVar)
	return q
}
