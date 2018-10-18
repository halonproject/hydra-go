#!/usr/bin/env bash

export ERROR_LOG="/tmp/error.log"
ERROR_MSG=$(cat $ERROR_LOG)

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
    GITHUB_API_URL="https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/issues/$CIRCLE_PR_NUMBER/comments"

    echo "posting error message to $GITHUB_API_URL"

    ESCAPED_ERROR_MSG=$(echo $ERROR_MSG | sed 's/"/\\\"/g')

    curl --request POST \
  --url $GITHUB_API_URL \
  --header 'accept: application/vnd.github.v3+json' \
  --header 'content-type: application/json' \
  -u cpurta:$GITHUB_ACCESS_TOKEN \
  --data "{\"body\": \"There was an error during the CI process:\n\`\`\`\n$ESCAPED_ERROR_MSG\n\`\`\`\nPlease check that your changes are working as intended.\"}"
}
