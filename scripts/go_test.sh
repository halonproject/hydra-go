#!/usr/bin/env bash

. scripts/report_error.sh

set -e

go test -v ./... | tee -a $ERROR_LOG >&2
