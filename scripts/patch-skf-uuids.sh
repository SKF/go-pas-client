#!/usr/bin/env bash

set -o pipefail
set -o nounset

NEEDLE='"format": "uuid"'
read -r -d '' REPLACE << EOM
"x-go-type": {
  "type": "UUID",
  "import": {
    "package": "github.com/SKF/go-utility/v2/uuid"
  },
  "hints": {
    "kind": "object"
  }
},
"format": "uuid"
EOM

set -o errexit

# Escape replacement for sed
REPLACE=$(sed 's/[&/\]/\\&/g' <<< "$(echo ${REPLACE})")

sed -i "s/${NEEDLE}/${REPLACE}/g" $1
