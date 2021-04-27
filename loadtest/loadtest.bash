#!/bin/bash
timecmd='/usr/bin/time -v'
merge_tool='./interval_merger'
interval_ranges='5 50 100 500 1000 5000 10000 50000 100000 500000 1000000'

for i in ${interval_ranges}
do
  input_file="${i}_intervals.json"
  CMD="${timecmd} ${merge_tool} --file=${input_file}"

  echo ${CMD}
  ${CMD}
  echo "---------------------------------------------------------"
done
