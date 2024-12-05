#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <day number>"
    exit 1
fi

DAY_NUM=$1
DIR_NAME="day$DAY_NUM"

mkdir -p $DIR_NAME
cd $DIR_NAME

go mod init $DIR_NAME

mkdir -p utils
cat > utils/utils.go << 'EOF'
package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ParseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var content []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return content, nil
}
EOF

for part in part1 part2; do
    mkdir -p $part
    cat > $part/main.go << EOF
package main

import (
	"$DIR_NAME/utils"
	"fmt"
)

func main() {
	content, err := utils.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}
EOF
done

echo "Project structure created successfully in $DIR_NAME"