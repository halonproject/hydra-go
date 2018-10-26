#!/usr/bin/env bash

. scripts/report_error.sh

set -e

go test -cover -coverprofile=coverage.txt 2>$ERROR_LOG
go get github.com/mattn/goveralls 2>$ERROR_LOG
goveralls -service=circleci -coverprofile=coverage.txt -repotoken=$COVERALLS_REPO_TOKEN 2>$ERROR_LOG
