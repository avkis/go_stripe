#! /bin/bash

if [ ! -f .env ]; then
	export $(cat .env | xargs)
fi

set -a && source .env && set +a
