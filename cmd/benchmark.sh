#!/bin/bash

AKUMULI_1host12hr=akumuli_queries_devops_scale8_1host12hr_month.gz
AKUMULI_1host1hr=akumuli_queries_devops_scale8_1host1hr_month.gz
AKUMULI_8host1hr=akumuli_queries_devops_scale8_8host1hr_month.gz
AKUMULI_8host1hr=akumuli_queries_devops_scale8_1host1day_month.gz

INFLUX_1host12hr=influx_queries_devops_scale8_1host12hr_month.gz
INFLUX_1host1hr=influx_queries_devops_scale8_1host1hr_month.gz
INFLUX_8host1hr=influx_queries_devops_scale8_8host1hr_month.gz
INFLUX_8host1hr=influx_queries_devops_scale8_1host1day_month.gz

echo ""
echo "Benchmarking Akumuli"

cat $AKUMULI_1host12hr | gunzip | ./query_benchmarker_akumuli/query_benchmarker_akumuli --urls=http://localhost:8181 --print-interval=0 --limit=1000 --workers=2
cat $AKUMULI_1host1hr | gunzip | ./query_benchmarker_akumuli/query_benchmarker_akumuli --urls=http://localhost:8181 --print-interval=0 --limit=1000 --workers=2
cat $AKUMULI_8host1hr | gunzip | ./query_benchmarker_akumuli/query_benchmarker_akumuli --urls=http://localhost:8181 --print-interval=0 --limit=1000 --workers=2
cat $AKUMULI_1host1day | gunzip | ./query_benchmarker_akumuli/query_benchmarker_akumuli --urls=http://localhost:8181 --print-interval=0 --limit=1000 --workers=2

echo ""
echo "Benchmarking Influx"

cat $INFLUX_1host12hr | gunzip | ./query_benchmarker_influxdb/query_benchmarker_influxdb --url=http://localhost:8086 --print-interval=0 --limit=1000 --workers=2
cat $INFLUX_1host1hr | gunzip | ./query_benchmarker_influxdb/query_benchmarker_influxdb --url=http://localhost:8086 --print-interval=0 --limit=1000 --workers=2
cat $INFLUX_8host1hr | gunzip | ./query_benchmarker_influxdb/query_benchmarker_influxdb --url=http://localhost:8086 --print-interval=0 --limit=1000 --workers=2
cat $INFLUX_1host1day | gunzip | ./query_benchmarker_influxdb/query_benchmarker_influxdb --url=http://localhost:8086 --print-interval=0 --limit=1000 --workers=2
