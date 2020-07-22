#!/usr/bin/env bash

# Filter markets for MARKET, LATDEG,LATMIN,LATSEC,LNGDEG,LNGMIN,LNGSEC,LATHEM,LNGHEM,NATION,DESCRIPTION
INPUT="$1"
[[ -z "$INPUT" ]] && echo "Need input file. Exit now" && exit 1
cat "$INPUT" | cut -d ',' -f1,8,9,10,11,12,13,18,19,26,29