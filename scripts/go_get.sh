#!/usr/bin/env bash

. scripts/report_error.sh

trap report_error_to_github EXIT

set -e

go get -v -t -d ./... 2>$ERROR_LOG
