#!/bin/bash

# Load data

AKUMULI_DEVOPS_GZ=akumuli_devops_scalevar8_1month.gz
INFLUX_DEVOPS_GZ=influx_devops_scalevar8_1month.gz

echo "Loading data into Akumuli"

cat $AKUMULI_DEVOPS_GZ | gunzip | ./bulk_load_akumuli/bulk_load_akumuli --batch-size=5000 --workers=1

echo "Loading data into InfluxDB"

cat $INFLUX_DEVOPS_GZ | gunzip | ./bulk_load_influx/bulk_load_influx --batch-size=5000 --workers=1
