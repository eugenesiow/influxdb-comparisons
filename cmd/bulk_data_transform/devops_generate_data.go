package main

import (
	"time"
	"encoding/csv"
	"os"
	"log"
	"io"
	"strconv"
	"fmt"
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

	GreenTaxiFields  = [][]byte{
		[]byte("vendorid"),
		[]byte("lpep_dropoff_datetime"),
		[]byte("store_and_fwd_flag"),
		[]byte("ratecodeid"),
		[]byte("pickup_longitude"),
		[]byte("pickup_latitude"),
		[]byte("dropoff_longitude"),
		[]byte("dropoff_latitude"),
		[]byte("passenger_count"),
		[]byte("trip_distance"),
		[]byte("fare_amount"),
		[]byte("extra"),
		[]byte("mta_tax"),
		[]byte("tip_amount"),
		[]byte("tolls_amount"),
		[]byte("improvement_surcharge"),
		[]byte("total_amount"),
		[]byte("payment_type"),
		[]byte("trip_type"),
		[]byte("point_date"),
	}
)

// A DevopsSimulator generates data similar to telemetry from Telegraf.
// It fulfills the Simulator interface.
type DevopsSimulator struct {
	eof bool
	reader *csv.Reader
	useCase string
}

func (g *DevopsSimulator) Finished() bool {
	//g.fileObj.Close()
	return g.eof
}

// Type DevopsSimulatorConfig is used to create a DevopsSimulator.
type DevopsSimulatorConfig struct {
	filePath string
	useCase string
}

func (d *DevopsSimulatorConfig) ToSimulator() *DevopsSimulator {
	file, err := os.Open(d.filePath)
	if err != nil {
		log.Fatal(err)
	}
	dg := &DevopsSimulator{
		reader: csv.NewReader(file),
		eof: false,
		useCase: useCase,
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
	if useCase=="shelburne" {
		if record[0] == "Time" {
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
			msFloat, _ := strconv.ParseFloat(record[i+1],64)
			p.AppendField(ShelburneFields[i], msFloat)
		}
	} else if useCase=="green_taxi" {
		if record[0] == "vendorid" {
			d.Next(p)
			return
		}

		p.SetMeasurementName([]byte("green_taxi"))
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, record[1])
		if err != nil {
			fmt.Println(err)
			fmt.Println(record)
		}
		p.SetTimestamp(&t)

		msLong, _ := strconv.ParseInt(record[0], 10,64)
		p.AppendField(GreenTaxiFields[0],msLong)
		p.AppendField(GreenTaxiFields[1],record[2])
		if record[3]=="true" {
			p.AppendField(GreenTaxiFields[2],true)
		} else {
			p.AppendField(GreenTaxiFields[2],false)
		}
		msLong2, _ := strconv.ParseInt(record[4], 10,64)
		p.AppendField(GreenTaxiFields[3],msLong2)
		msDouble, _ := strconv.ParseFloat(record[5],64)
		p.AppendField(GreenTaxiFields[4],msDouble)
		msDouble1, _ := strconv.ParseFloat(record[6],64)
		p.AppendField(GreenTaxiFields[5],msDouble1)
		msDouble2, _ := strconv.ParseFloat(record[7],64)
		p.AppendField(GreenTaxiFields[6],msDouble2)
		msDouble3, _ := strconv.ParseFloat(record[8],64)
		p.AppendField(GreenTaxiFields[7],msDouble3)
		msLong3, _ := strconv.ParseInt(record[9], 10,64)
		p.AppendField(GreenTaxiFields[8],msLong3)
		msDouble4, _ := strconv.ParseFloat(record[10],64)
		p.AppendField(GreenTaxiFields[9],msDouble4)
		msDouble5, _ := strconv.ParseFloat(record[11],64)
		p.AppendField(GreenTaxiFields[10],msDouble5)
		msDouble6, _ := strconv.ParseFloat(record[12],64)
		p.AppendField(GreenTaxiFields[11],msDouble6)
		msDouble7, _ := strconv.ParseFloat(record[13],64)
		p.AppendField(GreenTaxiFields[12],msDouble7)
		msDouble8, _ := strconv.ParseFloat(record[14],64)
		p.AppendField(GreenTaxiFields[13],msDouble8)
		msDouble9, _ := strconv.ParseFloat(record[15],64)
		p.AppendField(GreenTaxiFields[14],msDouble9)
		msDouble10, _ := strconv.ParseFloat(record[17],64)
		p.AppendField(GreenTaxiFields[15],msDouble10)
		msDouble11, _ := strconv.ParseFloat(record[18],64)
		p.AppendField(GreenTaxiFields[16],msDouble11)
		msLong4, _ := strconv.ParseInt(record[19], 10,64)
		p.AppendField(GreenTaxiFields[17],msLong4)
		if record[20]=="" {
			p.AppendField(GreenTaxiFields[18],1)
		} else {
			msLong5, _ := strconv.ParseInt(record[20], 10,64)
			p.AppendField(GreenTaxiFields[18],msLong5)
		}
		p.AppendField(GreenTaxiFields[19],record[21])

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