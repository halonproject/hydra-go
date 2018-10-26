#!/usr/bin/env bash

function diff_contains_error_line {
    HUNK_HEADER_LINE=0
    HUNK_HEADER_FOUND='false'
    I=1
    DIFF_POSITION=0

    ERROR_FILE_FOUND='false'
    ERROR_LINE_FOUND='false'

    DIFF=$1
    FILENAME=$2
    LINE_NUMBER=$3

    while read LINE; do
        case "$LINE" in
            diff* )
                continue
                ;;
            ---* )
                FROM_FILE=($LINE)
                if [[ ${FROM_FILE[1]} == *$FILENAME* ]]; then
                    ERROR_FILE_FOUND='true'
                fi
                ;;
            +++* )
                TO_FILE=($LINE)
                if [[ ${TO_FILE[1]} == *$FILENAME* ]]; then
                    ERROR_FILE_FOUND='true'
                fi
                ;;
            @@*@@* )
                if [[ $HUNK_HEADER_FOUND -eq "false" ]]; then
                    HUNK_HEADER_FOUND='true'
                    HUNK_HEADER_LINE=$I
                fi

                FILE_RANGES=($LINE)
                FROM_FILE_RANGE=${LINE[1]}
                TO_FILE_RANGE=${LINE[2]}

                # if the error line is between the start-from-range + num_lines; then DIFF_POSITION = $DIFF_POSITION + ($START_FROM_FILE_LINE - $LINE_NUMBER); break; fi

                # if the error line is between the start-from-range + num_lines; then DIFF_POSITION = $DIFF_POSITION + ($START_TO_FILE_LINE - $LINE_NUMBER); break; fi
                ;;
            * )
                continue
                ;;
        esac
        let "I = $I + 1"
        if [[ $HUNK_HEADER_FOUND -eq "true" ]]; then
            let "DIFF_POSITION = $I - $HUNK_HEADER_LINE"
        fi
    done < $DIFF

    return DIFF_POSITION
}
