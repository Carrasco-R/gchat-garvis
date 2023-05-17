#!/usr/bin/env bash

while getopts ':w:r:' OPTION; do
  case "$OPTION" in
    w)
      WHO=${OPTARG}
      ;;
    r)
      REASON=${OPTARG}
      ;;
    ?)
      echo "script usage: $(basename \$0) [-w who] [-r reason]" >&2
      exit 1
      ;;
  esac
done
shift "$(($OPTIND -1))"

echo "\nRecognition to ${WHO} ${REASON}\n"
read -r -p "Are you sure you want to send? [y/n] " response
if [ "$response" = "y" ]; then
  curl --request POST --url ${CONTAINER_IP}/recognize --header "content-type: application/json" --data '{"person":"'"$WHO"'","reason":"'"$REASON"'"}'
fi