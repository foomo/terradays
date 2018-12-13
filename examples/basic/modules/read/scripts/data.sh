#!/bin/bash

# Parse arguements
eval "$(jq -r '@sh "FILENAME=\(.filename)"')"


# Validate file exists
if [ ! -e $FILENAME ]; then
  echo "Could not read ${FILENAME}" >&2
  exit 1
fi

# Read in data
cat $FILENAME
