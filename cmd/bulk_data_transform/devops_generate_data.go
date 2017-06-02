package main

import (
	"time"
	"encoding/csv"
	"os"
	"log"
	"io"
	"strconv"
	"fmt"
	"path"
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

	GreenTaxiCols = []int{
		0,2,3,4,5,6,7,8,9,10,11,12,13,14,15,17,18,19,20,21,
	}

	numOfCols = 0
)
var srBenchFields []string

// A DevopsSimulator generates data similar to telemetry from Telegraf.
// It fulfills the Simulator interface.
type DevopsSimulator struct {
	eof bool
	reader *csv.Reader
	useCase string
	format string
	filePath string
}

func (g *DevopsSimulator) Finished() bool {
	//g.fileObj.Close()
	return g.eof
}

// Type DevopsSimulatorConfig is used to create a DevopsSimulator.
type DevopsSimulatorConfig struct {
	filePath string
	useCase string
	format string
}

func (d *DevopsSimulatorConfig) ToSimulator() *DevopsSimulator {
	file, err := os.Open(d.filePath)
	if err != nil {
		log.Fatal(err)
	}
	dg := &DevopsSimulator{
		reader: csv.NewReader(file),
		eof: false,
		useCase: d.useCase,
		format: d.format,
		filePath: d.filePath,
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
		if format=="akumuli" {
			p.AppendTag([]byte("shelburne"),[]byte("wsda_sensor"))
		}

		for i := range ShelburneFields {
			if record[i+1] == "" || record[i+1] == "NaN" {
				p.Reset()
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

		//for i := range GreenTaxiFields {
		//	if record[GreenTaxiCols[i]] == "" || record[GreenTaxiCols[i]] == "NaN" {
		//		d.Next(p)
		//		return
		//	}
		//	p.AppendField(GreenTaxiFields[i], record[GreenTaxiCols[i]])
		//}

		msLong, _ := strconv.ParseInt(record[0], 10,64)
		p.AppendField(GreenTaxiFields[0],msLong)
		t2, err := time.Parse(layout, record[2])
		if err != nil {
			fmt.Println(err)
			fmt.Println(record)
		}
		p.AppendField(GreenTaxiFields[1],t2.Nanosecond())
		p.AppendTag([]byte("taxi"),[]byte("green"))
		if record[3]=="true" {
			p.AppendField(GreenTaxiFields[2],1)
		} else {
			p.AppendField(GreenTaxiFields[2],0)
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
		t3, err := time.Parse(layout, record[21])
		if err != nil {
			fmt.Println(err)
			fmt.Println(record)
		}
		p.AppendField(GreenTaxiFields[19],t3.Nanosecond())

	} else if useCase=="srbench" {
		if record[0] == "time" {
			numOfCols = len(record)
			if numOfCols<=1 {
				p.Reset()
				return
			}
			srBenchFields = make([]string, numOfCols)
			copy(srBenchFields, record)
			d.Next(p)
			return
		}
		fName := path.Base(d.filePath)
		extName := path.Ext(d.filePath)
		bName := fName[:len(fName)-len(extName)]

		p.SetMeasurementName([]byte("_"+bName))
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, record[0])
		if err != nil {
			fmt.Println(err)
			fmt.Println(record)
		}
		p.SetTimestamp(&t)
		if t.Nanosecond() < 0 {
			p.Reset()
			return
		}

		if format=="akumuli" {
			p.AppendTag([]byte("srbench"),[]byte("lsd_blizzard"))
		}

		for i := 0; i < numOfCols - 1; i++ {
			if record[i+1]=="false" {
				p.AppendField([]byte(srBenchFields[i+1]), 0.0)
			} else if record[i+1]=="true" {
				p.AppendField([]byte(srBenchFields[i+1]), 1.0)
			} else {
				msFloat, _ := strconv.ParseFloat(record[i+1],64)
				p.AppendField([]byte(srBenchFields[i+1]), msFloat)
			}
		}
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