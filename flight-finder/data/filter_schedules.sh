#!/usr/bin/env bash

# Filter schedules for FROM,TO,CARRIER
INPUT="$1"
[[ -z "$INPUT" ]] && echo "Need input file. Exit now" && exit 1
cat "$INPUT" | cut -d ',' -f1,2,4 | sort | uniq