package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"text/template"
	"time"
)

// AkumuliDevops produces Akumuli-specific queries for all the devops query types.
type AkumuliDevops struct {
	AllInterval TimeInterval
}

// NewAkumuliDevops makes an AkumuliDevops object ready to generate Queries.
func newAkumuliDevopsCommon(dbConfig DatabaseConfig, start, end time.Time) QueryGenerator {
	if !start.Before(end) {
		panic("bad time order")
	}

	return &AkumuliDevops{
		AllInterval: NewTimeInterval(start, end),
	}
}

// Dispatch fulfills the QueryGenerator interface.
func (d *AkumuliDevops) Dispatch(i, scaleVar int) Query {
	q := NewHTTPQuery() // from pool
	devopsDispatchAll(d, i, q, scaleVar)
	return q
}

func (d *AkumuliDevops) MaxCPUUsageHourByMinuteOneHost(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 1, time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsageHourByMinuteTwoHosts(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 2, time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsageHourByMinuteFourHosts(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 4, time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsageHourByMinuteEightHosts(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 8, time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsageHourByMinuteSixteenHosts(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 16, time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsageHourByMinuteThirtyTwoHosts(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 32, time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsage12HoursByMinuteOneHost(q Query, scaleVar int) {
	d.maxCPUUsageHourByMinuteNHosts(q.(*HTTPQuery), scaleVar, 1, 12*time.Hour)
}

func (d *AkumuliDevops) MaxCPUUsageDayByHour(q Query, scaleVar int) {
	d.maxCPUUsageDayByHourNHosts(q.(*HTTPQuery), scaleVar, 1, 12*time.Hour)
}

func (d *AkumuliDevops) maxCPUUsageDayByHourNHosts(qi Query, scaleVar, nhosts int, timeRange time.Duration) {
	interval := d.AllInterval.RandWindow(timeRange)
	nn := rand.Perm(scaleVar)[:nhosts]
	hostnames := []string{}
	for _, n := range nn {
		hostnames = append(hostnames, fmt.Sprintf("\"host_%d\"", n))
	}

	combinedHostnameClause := strings.Join(hostnames, ",")

	startTimestamp := interval.StartUnixNano()
	endTimestamp := interval.EndUnixNano()

	const tmplString = `
	{
		"group-aggregate": {
			"metric": "cpu.usage_user",
			"func": [ "max" ],
			"step": "1h"
		},
		"range": {
			"from": {{.StartTimestamp}},
			"to": {{.EndTimestamp}}
		},
		"where": {
			"hostname": [ {{.CombinedHostnameClause}} ]
		},
		"output": {
			"format": "csv"
		}
	}
	`

	tmpl := template.Must(template.New("tmpl").Parse(tmplString))
	bodyWriter := new(bytes.Buffer)

	arg := struct {
		StartTimestamp, EndTimestamp int64
		CombinedHostnameClause       string
	}{
		startTimestamp,
		endTimestamp,
		combinedHostnameClause,
	}
	err := tmpl.Execute(bodyWriter, arg)
	if err != nil {
		panic("logic error")
	}

	humanLabel := fmt.Sprintf("Akumuli max cpu, rand %4d hosts, rand %s by 1m", nhosts, timeRange)
	q := qi.(*HTTPQuery)
	q.HumanLabel = []byte(humanLabel)
	q.HumanDescription = []byte(fmt.Sprintf("%s: %s", humanLabel, interval.StartString()))
	q.Method = []byte("POST")
	q.Path = []byte("/api/query")
	q.Body = bodyWriter.Bytes()
	q.StartTimestamp = interval.StartUnixNano()
	q.EndTimestamp = interval.EndUnixNano()
}

func (d *AkumuliDevops) maxCPUUsageHourByMinuteNHosts(qi Query, scaleVar, nhosts int, timeRange time.Duration) {
	interval := d.AllInterval.RandWindow(timeRange)
	nn := rand.Perm(scaleVar)[:nhosts]

	hostnames := []string{}
	for _, n := range nn {
		hostnames = append(hostnames, fmt.Sprintf("\"host_%d\"", n))
	}

	combinedHostnameClause := strings.Join(hostnames, ",")

	startTimestamp := interval.StartUnixNano()
	endTimestamp := interval.EndUnixNano()

	const tmplString = `
	{
		"group-aggregate": {
			"metric": "cpu.usage_user",
			"func": [ "max" ],
			"step": "1m"
		},
		"range": {
			"from": {{.StartTimestamp}},
			"to": {{.EndTimestamp}}
		},
		"where": {
			"hostname": [ {{.CombinedHostnameClause}} ]
		},
		"output": {
			"format": "csv"
		}
	}
	`

	tmpl := template.Must(template.New("tmpl").Parse(tmplString))
	bodyWriter := new(bytes.Buffer)

	arg := struct {
		StartTimestamp, EndTimestamp int64
		CombinedHostnameClause       string
	}{
		startTimestamp,
		endTimestamp,
		combinedHostnameClause,
	}
	err := tmpl.Execute(bodyWriter, arg)
	if err != nil {
		panic("logic error")
	}

	humanLabel := fmt.Sprintf("Akumuli max cpu, rand %4d hosts, rand %s by 1m", nhosts, timeRange)
	q := qi.(*HTTPQuery)
	q.HumanLabel = []byte(humanLabel)
	q.HumanDescription = []byte(fmt.Sprintf("%s: %s", humanLabel, interval.StartString()))
	q.Method = []byte("POST")
	q.Path = []byte("/api/query")
	q.Body = bodyWriter.Bytes()
	q.StartTimestamp = interval.StartUnixNano()
	q.EndTimestamp = interval.EndUnixNano()
}

func (d *AkumuliDevops) MeanCPUUsageDayByHourAllHostsGroupbyHost(qi Query, _ int) {
	interval := d.AllInterval.RandWindow(24 * time.Hour)

	startTimestamp := interval.StartUnixNano()
	endTimestamp := interval.EndUnixNano()

	const tmplString = `
	{
		"group-aggregate": {
			"metric": "cpu.usage_user",
			"func": [ "avg" ],
			"step": "1h"
		},
		"range": {
			"from": {{.StartTimestamp}},
			"to": {{.EndTimestamp}}
		},
		"group-by": {
			"tag": "hostname"
		},
		"output": {
			"format": "csv"
		}
	}
	`

	tmpl := template.Must(template.New("tmpl").Parse(tmplString))
	bodyWriter := new(bytes.Buffer)

	arg := struct {
		StartTimestamp, EndTimestamp int64
	}{
		startTimestamp,
		endTimestamp,
	}
	err := tmpl.Execute(bodyWriter, arg)
	if err != nil {
		panic("logic error")
	}

	humanLabel := fmt.Sprintf("Akumuli mean cpu, all hosts, rand %s by 1h", interval.Duration().String())
	q := qi.(*HTTPQuery)
	q.HumanLabel = []byte(humanLabel)
	q.HumanDescription = []byte(fmt.Sprintf("%s: %s", humanLabel, interval.StartString()))
	q.Method = []byte("POST")
	q.Path = []byte("/api/query")
	q.Body = bodyWriter.Bytes()
	q.StartTimestamp = interval.StartUnixNano()
	q.EndTimestamp = interval.EndUnixNano()
}

