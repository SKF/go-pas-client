#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

OUTPUT=$(
  cat "$1" \
  | jq '.definitions.Node["x-nullable"] = false' \
  | jq '.definitions.NodeMetaData.additionalProperties["x-nullable"] = true' \
)

[[ $? == 0 ]] && echo "${OUTPUT}" >| $1
