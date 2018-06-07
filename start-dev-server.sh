#!/bin/sh

function restart {
	printf "\033c"
	go build -o build/fc
	./build/fc &
}

function killExec {
	kill $!
}

trap killExec EXIT

restart

while inotifywait -qq -r -e modify . @./.git @./.idea @./vendor @./build; do
	killExec
	restart
done
