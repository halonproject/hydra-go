#!/usr/bin/env bash

PROJECT_OWNER="halonproject"
ERROR_MSG=""
FILEPATH=""
export ERROR_LOG="/tmp/error.log"

function report_error_to_github {
    if [ $? -eq 0 ]; then
        exit 0
    fi

    ERROR_MSG=$(cat $ERROR_LOG)

    echo $ERROR_MSG

    printenv

    post_error_to_github
}

function get_pull_request_diff {
    curl --request GET \
  --url https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/pulls/$CIRCLE_PR_NUMBER \
  --header 'accept: application/vnd.github.v3.diff'
}

function post_error_to_github {
    echo "posting error message to https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/pulls/$CIRCLE_PR_NUMBER/comments"

    curl --request POST \
  --url https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/pulls/$CIRCLE_PR_NUMBER/comments \
  --header 'accept: application/vnd.github.v3+json' \
  --header 'content-type: application/json' \
  --data "{
	\"body\": \"$ERROR_MSG\",
	\"commit_id\": \"$CIRCLE_SHA1\",
	\"path\": \"\",
	\"position\": 0
}"
}
