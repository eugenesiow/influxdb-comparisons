#!/bin/bash

free
echo "Drop caches"
echo 1 >/proc/sys/vm/drop_caches
echo 2 >/proc/sys/vm/drop_caches
echo 3 >/proc/sys/vm/drop_caches
echo "Done"
free
