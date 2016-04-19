#!/bin/bash
CMD="$@"
if [ -z "$CMD" ]; then
	echo "You must specify a command to execute."
else
	"$CMD" > /dev/null 2>&1 &
fi

