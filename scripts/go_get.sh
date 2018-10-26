#!/usr/bin/env bash

. scripts/report_error.sh

set -e

go get -v -t -d ./... | tee -a $ERROR_LOG >&2
