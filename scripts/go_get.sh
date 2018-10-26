#!/usr/bin/env bash

. scripts/report_error.sh

set -e

go get -v -t -d ./... 2>$ERROR_LOG
