#!/bin/sh

EXCLUDES="-e './.git' -e './.idea' -e './vendor' -e './build'"
EXEC_NAME="fleet-commander___"

function killExec {
	kill `pgrep -f "$EXEC_NAME"`
}

function start {
	printf "\033c"
	vgo build -o build/$EXEC_NAME
	./build/$EXEC_NAME &	
}

start

fswatch -i 0,2 -o --event=Updated -e $EXCLUDES ./ | while read; do
	killExec
	start
done

killExec
