#!/usr/bin/env bash

export ERROR_LOG="/tmp/error.log"

function report_error_to_github {
    if [ $? -eq 0 ]; then
        exit 0
    fi

    ERROR_MSG=$(cat $ERROR_LOG)

    echo $ERROR_MSG

    FILES_FOUND=$(echo $ERROR_MSG | grep "[\w-]+\.[A-Za-z]{1,3}")

    if [ ${#FILES_FOUND} -ne 0 ]; then
        echo "Errors found in files: ${FILES_FOUND[@]}"
    fi

    if [[ ! -z $CI && $CI == "true" ]]; then
        post_pull_request_comment
    fi
}

function post_pull_request_diff_comment {
    if [ -z $CIRCLE_PR_NUMBER ]; then
        CIRCLE_PR_NUMBER="${CIRCLE_PULL_REQUEST//[^0-9]/}"
    fi

    GITHUB_API_URL="https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/pulls/$CIRCLE_PR_NUMBER/comments"

    echo "posting error message to $GITHUB_API_URL"

    ESCAPED_ERROR_MSG=$(echo $ERROR_MSG | sed 's/"/\\\"/g')

    curl --request POST \
  --url $GITHUB_API_URL \
  --header 'accept: application/vnd.github.v3+json' \
  --header 'content-type: application/json' \
  -u cpurta:$GITHUB_ACCESS_TOKEN \
  --data "{\"body\": \"There was an error during the CI process:\n\`\`\`\n$ESCAPED_ERROR_MSG\n\`\`\`\nPlease check that your changes are working as intended.\"}"
}

function post_pull_request_comment {
    if [ -z $CIRCLE_PR_NUMBER ]; then
        CIRCLE_PR_NUMBER="${CIRCLE_PULL_REQUEST//[^0-9]/}"
    fi

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

trap report_error_to_github EXIT
