#!/bin/bash

run_parts() {
    local day=$1
    cd $day || return
    
    for part in part*; do
        if [ -d "$part" ] && [ -f "$part/main.go" ]; then
            output=$(go run $part/main.go)
            echo "$day, $part = $output"
        fi
    done
    
    cd ..
    echo
}

if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed or not in PATH"
    exit 1
fi

if [ $# -eq 1 ]; then
    # Run specific day
    day="day$1"
    if [ -d "$day" ]; then
        run_parts $day
    else
        echo "Day $1 not found"
        exit 1
    fi
else
    # Run all days
    for day in day*; do
        if [ -d "$day" ]; then
            run_parts $day
        fi
    done
fi