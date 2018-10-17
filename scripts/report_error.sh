#!/usr/bin/env bash

PROJECT_OWNER="halonproject"
FILEPATH=""
ERROR_MSG=$(cat $ERROR_LOG)
export ERROR_LOG="/tmp/error.log"

function report_error_to_github {
    if [ $? -eq 0 ]; then
        exit 0
    fi

    echo $ERROR_MSG

    if [ -z $CIRCLE_PR_NUMBER ]; then
        CIRCLE_PR_NUMBER="${CIRCLE_PULL_REQUEST//[^0-9]/}"
    fi

    post_error_to_github
}

function get_pull_request_diff {
    curl --request GET \
  --url https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/pulls/$CIRCLE_PR_NUMBER \
  --header 'accept: application/vnd.github.v3.diff'
}

function post_error_to_github {
    echo "posting error message to https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/issues/$CIRCLE_PR_NUMBER/comments"

    curl --request POST \
  --url https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/issues/$CIRCLE_PR_NUMBER/comments \
  --header 'accept: application/vnd.github.v3+json' \
  --header 'content-type: application/json' \
  -u cpurta:$GITHUB_ACCESS_TOKEN \
  --data "{\"body\": \"There was an error during the CI process:\n$ERROR_MSG\nPlease check that your changes are working as intended.\"}"
}
