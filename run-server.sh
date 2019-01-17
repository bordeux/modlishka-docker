#!/bin/sh
set -e

RUN_CMD='Modlishka'
env | while IFS='=' read -r -d '' envName envValue; do
	envName=$(echo "${envName}" | tr '[:upper:]' '[:lower:]')
	if [[ ${envName} == ml* ]] ;
	then
		OPTION_NAME=$(echo "${envName}" | sed -r 's/(^ml\_)//g' | sed -r 's/(_)([a-z])/\U\2/g')
		RUN_CMD="${RUN_CMD} -${OPTION_NAME} ${envValue}"
	fi
done


${RUN_CMD}