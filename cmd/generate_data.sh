#!/bin/bash

# Generate data

echo "Generate data for Akumuli"
./bulk_data_gen/bulk_data_gen --seed=123 --use-case=devops --scale-var=8 --format=akumuli --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z | gzip > akumuli_devops_scalevar8_1month.gz

echo "Generate data for InfluxDB"
./bulk_data_gen/bulk_data_gen --seed=123 --use-case=devops --scale-var=8 --format=influx-bulk --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z | gzip > influx_devops_scalevar8_1month.gz

echo "Generate queries for Akumuli"
./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=akumuli --use-case=devops --query-type=1-host-1-hr --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > akumuli_queries_devops_scale8_1host1hr_month.gz

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=akumuli --use-case=devops --query-type=8-host-1-hr --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > akumuli_queries_devops_scale8_8host1hr_month.gz

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=akumuli --use-case=devops --query-type=1-host-12-hr --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > akumuli_queries_devops_scale8_1host12hr_month.gz

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=akumuli --use-case=devops --query-type=1-host-1-day --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > akumuli_queries_devops_scale8_1host1day_month.gz

echo "Generate queries for InfluxDB"

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=influx-http --use-case=devops --query-type=1-host-1-hr --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > influx_queries_devops_scale8_1host1hr_month.gz

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=influx-http --use-case=devops --query-type=8-host-1-hr --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > influx_queries_devops_scale8_8host1hr_month.gz

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=influx-http --use-case=devops --query-type=1-host-12-hr --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > influx_queries_devops_scale8_1host12hr_month.gz

./bulk_query_gen/bulk_query_gen --debug=0 --seed=123 --format=influx-http --use-case=devops --query-type=1-host-1-day --scale-var=8 --timestamp-start=2016-01-01T00:00:00Z --timestamp-end=2016-02-01T00:00:00Z  | gzip > influx_queries_devops_scale8_1host1day_month.gz
