#!/bin/bash

# disable cgo checks. But it doesnt help with "panic: runtime error: cgo argument has Go pointer to Go pointer"
#export GODEBUG=cgocheck=0

for dict in /usr/share/dist/words /usr/share/cracklib/cracklib-small
do
  if [ -f "$dict" ]; then
      break
  fi
  echo "Dictionary: $dict doesnt exist"
done

if [ -z "$dict" ]; then
    echo "No dictionary to compress found. Exit now"
    exit 1
fi

echo "Selected dictionary: $dict"
echo "Uncompressed size vs compressed size:"
cat $dict | wc -c 
cat $dict | go run . | wc -c

