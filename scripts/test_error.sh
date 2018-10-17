#!/usr/bin/env bash

. scripts/report_error.sh

trap report_error_to_github EXIT

set -e

rm -r ./temp 2>$ERROR_LOG

echo "we should not get to this line"
