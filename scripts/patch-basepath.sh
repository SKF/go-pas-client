#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

NEEDLE='"basePath": ""'
REPLACE='"basePath": "\/v1"'

sed -i "s/${NEEDLE}/${REPLACE}/g" $1
