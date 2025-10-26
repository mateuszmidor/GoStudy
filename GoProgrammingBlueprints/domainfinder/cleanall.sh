#!/bin/bash

# Remove all the tool binaries
tools=`find ../domainfinder/ -executable -type f | grep -v '.sh'`
echo "Removing:"
echo "$tools"
rm $tools
echo "Done."
