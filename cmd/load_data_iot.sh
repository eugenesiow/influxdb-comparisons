#!/bin/bash

# Load data

INFLUX_SHELBURNE_GZ=/scratch/influx_shelburne.gz

echo "Loading data into InfluxDB"

cat $INFLUX_SHELBURNE_GZ | gunzip | ./bulk_load_influx/bulk_load_influx --batch-size=250 --workers=1
