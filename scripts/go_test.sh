#!/usr/bin/env bash

. scripts/report_error.sh

set -e

go test -v ./... 2>$ERROR_LOG
