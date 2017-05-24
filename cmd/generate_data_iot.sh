echo "Generate data for InfluxDB"
./bulk_data_transform/bulk_data_transform --use-case=shelburne --format=influx-bulk --input="/scratch/shelburne.csv" | gzip > influx_shelburne.gz