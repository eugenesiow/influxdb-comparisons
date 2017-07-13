echo "Generate data for InfluxDB"
#./bulk_data_transform/bulk_data_transform --use-case=green_taxi --format=influx-bulk --input="/scratch/2016_green_taxi_trip_data.csv" | gzip > influx_green_taxi.gz
#./bulk_data_transform/bulk_data_transform --use-case=green_taxi --format=mongo --input="/scratch/2016_green_taxi_trip_data.csv" | gzip > mongo_green_taxi.gz
#./bulk_data_transform/bulk_data_transform --use-case=shelburne --format=mongo --input="/scratch/shelburne.csv" | gzip > mongo_shelburne.gz
#./bulk_data_transform/bulk_data_transform --use-case=srbench --format=mongo --input="/scratch/knoesis_observations_csv_date_sorted/" | gzip > mongo_srbench.gz
#./bulk_data_transform/bulk_data_transform --use-case=srbench --format=cassandra --input="/scratch/knoesis_observations_csv_date_sorted/" | gzip > cassandra_srbench.gz
#./bulk_data_transform/bulk_data_transform --use-case=shelburne --format=influx-bulk --input="/scratch/shelburne.csv" | gzip > influx_shelburne.gz
#./bulk_data_transform/bulk_data_transform --use-case=shelburne --format=cassandra --input="/scratch/shelburne.csv" | gzip > cassandra_shelburne.gz
./bulk_data_transform/bulk_data_transform --use-case=shelburne --format=opentsdb --input="/scratch/shelburne.csv" | gzip > opentsdb_shelburne.gz
#./bulk_data_transform/bulk_data_transform --use-case=srbench --format=influx-bulk --input="/scratch/knoesis_observations_csv_date_sorted/" | gzip > influx_srbench.gz
#./bulk_data_transform/bulk_data_transform --use-case=srbench --format=opentsdb --input="/scratch/knoesis_observations_csv_date_sorted/" | gzip > opentsdb_srbench.gz
#./bulk_data_transform/bulk_data_transform --use-case=green_taxi --format=cassandra --input="/scratch/2016_green_taxi_trip_data.csv" | gzip > cassandra_green_taxi.gz
#./bulk_data_transform/bulk_data_transform --use-case=green_taxi --format=opentsdb --input="/scratch/2016_green_taxi_trip_data.csv" | gzip > opentsdb_green_taxi.gz
