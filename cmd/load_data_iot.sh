#!/bin/bash

# Load data

INFLUX_SHELBURNE_GZ=/scratch/influx_shelburne.gz
INFLUX_SRBENCH_GZ=/scratch/influx_srbench.gz
INFLUX_GREEN_TAXI_GZ=/scratch/influx_green_taxi.gz
MONGO_GREEN_TAXI_GZ=/scratch/mongo_green_taxi.gz
MONGO_SHELBURNE_GZ=/scratch/mongo_shelburne.gz
MONGO_SRBENCH_GZ=/scratch/mongo_srbench.gz
CASSANDRA_SRBENCH_GZ=/scratch/cassandra_srbench.gz
CASSANDRA_SHELBURNE_GZ=/scratch/cassandra_shelburne.gz
CASSANDRA_GREEN_TAXI_GZ=/scratch/cassandra_green_taxi.gz
OPENTSDB_GREEN_TAXI_GZ=/scratch/opentsdb_green_taxi.gz
OPENTSDB_SRBENCH_GZ=/scratch/opentsdb_srbench.gz
OPENTSDB_SHELBURNE_GZ=/scratch/opentsdb_shelburne.gz


echo "Loading data into InfluxDB"

#cat $MONGO_GREEN_TAXI_GZ | gunzip | ./bulk_load_mongo/bulk_load_mongo --batch-size=1 --workers=4
#cat $MONGO_SHELBURNE_GZ | gunzip | ./bulk_load_mongo/bulk_load_mongo --batch-size=1 --workers=4
#cat $MONGO_SRBENCH_GZ | gunzip | ./bulk_load_mongo/bulk_load_mongo --batch-size=1 --workers=4
#cat $INFLUX_GREEN_TAXI_GZ | gunzip | ./bulk_load_influx/bulk_load_influx --batch-size=1 --workers=4
#cat $INFLUX_SRBENCH_GZ | gunzip | ./bulk_load_influx/bulk_load_influx --batch-size=1 --workers=4
cat $OPENTSDB_SHELBURNE_GZ | gunzip | ./bulk_load_opentsdb/bulk_load_opentsdb --batch-size=1 --workers=4
#cat $OPENTSDB_SRBENCH_GZ | gunzip | ./bulk_load_opentsdb/bulk_load_opentsdb --batch-size=1 --workers=4
#cat $CASSANDRA_SRBENCH_GZ | gunzip | ./bulk_load_cassandra/bulk_load_cassandra --batch-size=1 --workers=4
#cat $CASSANDRA_SHELBURNE_GZ | gunzip | ./bulk_load_cassandra/bulk_load_cassandra --batch-size=1 --workers=4
#cat $CASSANDRA_GREEN_TAXI_GZ | gunzip | ./bulk_load_cassandra/bulk_load_cassandra --batch-size=1 --workers=4
#cat $OPENTSDB_GREEN_TAXI_GZ | gunzip | ./bulk_load_opentsdb/bulk_load_opentsdb --batch-size=1 --workers=8
