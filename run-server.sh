#!/bin/sh
set -e

RUN_CMD='Modlishka'
IFS='
'
envList=$(env)
for line in $envList
do
        envName=$(echo "${line}" | cut -d'=' -f 1 | tr '[:upper:]' '[:lower:]')
        if [[ "${envName:0:3}" = "ml_" ]] ;
        then
			optionName=$(echo "${envName}" | sed -r 's/(^ml\_)//g' | sed -r 's/(_)([a-z])/\U\2/g')
			RUN_CMD="${RUN_CMD} -${optionName} ${optionName}"
        fi
done

echo "Running command: ${RUN_CMD}"

#${RUN_CMD}