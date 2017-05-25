echo "Generate data for InfluxDB"
./bulk_data_transform/bulk_data_transform --use-case=shelburne --format=influx-bulk --input="/scratch/shelburne.csv" | gzip > influx_shelburne.gz
./bulk_data_transform/bulk_data_transform --use-case=green_taxi --format=influx-bulk --input="/scratch/2016_green_taxi_trip_data.csv" | gzip > influx_green_taxi.gz