#!/bin/bash

set -e

day_num="$1"

cp -r day0/ day${day_num}/

sed -i "s/day0/day${day_num}/g" "day${day_num}/task.go"
sed -i "s/day0/day${day_num}/g" "day${day_num}/task_test.go"