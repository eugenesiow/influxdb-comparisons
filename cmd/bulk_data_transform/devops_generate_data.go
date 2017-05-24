package main

import (
	"time"
	"encoding/csv"
	"os"
	"log"
	"io"
	"strconv"
)

var (
	ShelburneFields = [][]byte{
		[]byte("air_temperature"),
		[]byte("solar_radiation"),
		[]byte("soil_moisture"),
		[]byte("leaf_wetness"),
		[]byte("internal_temperature"),
		[]byte("relative_humidity"),
	}
)

// A DevopsSimulator generates data similar to telemetry from Telegraf.
// It fulfills the Simulator interface.
type DevopsSimulator struct {
	eof bool
	reader *csv.Reader
	fileObj *os.File
}

func (g *DevopsSimulator) Finished() bool {
	g.fileObj.Close()
	return g.eof
}

// Type DevopsSimulatorConfig is used to create a DevopsSimulator.
type DevopsSimulatorConfig struct {
	filePath string
}

func (d *DevopsSimulatorConfig) ToSimulator() *DevopsSimulator {
	file, err := os.Open(d.filePath)
	if err != nil {
		log.Fatal(err)
	}
	dg := &DevopsSimulator{
		reader: csv.NewReader(file),
		eof: false,
		fileObj: file,
	}

	return dg
}

// Next advances a Point to the next state in the generator.
func (d *DevopsSimulator) Next(p *Point) {
	record, err := d.reader.Read()
	// Stop at EOF.
	if err == io.EOF {
		d.eof = true
		return
	}
	if record[0]=="Time" {
		d.Next(p)
		return
	}

	p.SetMeasurementName([]byte("wsda_sensor"))
	newTime := nsToTime(record[0])
	p.SetTimestamp(&newTime)

	for i := range ShelburneFields {
		if record[i+1] == "" || record[i+1] == "NaN" {
			d.Next(p)
			return
		}
		p.AppendField(ShelburneFields[i], record[i+1])
	}

	return
}

func nsToTime(ms string) (time.Time) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}
	}

	return time.Unix(0, msInt*int64(time.Nanosecond))
}