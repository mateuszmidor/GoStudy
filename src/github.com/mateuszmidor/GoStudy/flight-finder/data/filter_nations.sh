#!/usr/bin/env bash

# Filter nations for NATION,ISOALPHACODE,PRIMECUR,DESCRIPTION
# example: "PL","POL","PLN","POLAND"

INPUT="$1"
[[ -z "$INPUT" ]] && echo "Need input file. Exit now" && exit 1
cat "$INPUT" | cut -d ',' -f1,13,14,17