#!/bin/bash

export COMPOSE_PROJECT_NAME=localdev
local_cmdname=${0##*/}

usage()
{
	cat << USAGE >&2
Usage:
	$local_cmdname up/down
USAGE
	exit 1
}

up()
{
	docker-compose run --rm start_dependencies
}


# process arguments
while [[ $# -gt 0 ]]
do
	case "$1" in
		up)
		echo "Starting Containers"
		up
		break
		;;

		down)
		echo "Stopping Containers"
		down
		break
		;;

		*)
		echo "Unknown argument: $1" >&2
		usage
		;;
	esac
done

if [[ "$1" == "" ]]; then
	usage
fi
